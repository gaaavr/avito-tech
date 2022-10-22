package repository

import (
	"avito/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	tableUsers        = "users"
	columnBalance     = "balance"
	columnUserId      = "user_id"
	tableTransactions = "transactions"
	columnAmount      = "amount"
	columnDate        = "date_time"
	columnMessage     = "message"
)

var (
	errInsertRow = errors.New("failed to insert data into database")
)

// UserRepo - user object in the repository layer
type UserRepo struct {
	db *pgxpool.Pool
}

// NewUserRepo - constructor function for UserRepo
func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

// AccrualFunds - method of accruing cash to the balance
func (u *UserRepo) AccrualFunds(ac models.AccrualCash) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		return err
	}
	updateUserBalance := fmt.Sprintf("UPDATE %s SET %s=%s+$1 WHERE %s=$2",
		tableUsers, columnBalance, columnBalance, columnUserId)
	result, err := tx.Exec(context.Background(), updateUserBalance, ac.Amount, ac.UserID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		err = u.addUser(ac)
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}
	t := models.Transaction{
		UserID:  ac.UserID,
		Amount:  ac.Amount,
		Date:    time.Now().Truncate(time.Second),
		Message: ac.Message,
	}
	err = u.addTransaction(t)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

// BlockFunds - method of reserving funds from the main balance in a separate account
func (u *UserRepo) BlockFunds(order models.Order) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		return err
	}
	updateUserBalance := fmt.Sprintf("UPDATE %s SET %s=%s-$1 WHERE %s=$2",
		tableUsers, columnBalance, columnBalance, columnUserId)
	result, err := tx.Exec(context.Background(), updateUserBalance, order.Amount, order.UserID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback(context.Background())
		return fmt.Errorf("user with id %d does not exist", order.UserID)
	}

	createOrder := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4, $5, $6)", tableOrders)
	result, err = tx.Exec(context.Background(), createOrder, order.OrderID, order.UserID, order.ServiceID,
		order.Amount, order.Date, order.Block)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected = result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback(context.Background())
		return errInsertRow
	}

	t := models.Transaction{
		UserID:  order.UserID,
		Amount:  -order.Amount,
		Date:    order.Date,
		Message: "service payment",
	}
	err = u.addTransaction(t)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

// GetBalance - method to get user balance
func (u *UserRepo) GetBalance(ub *models.UserBalance) (*models.UserBalance, error) {
	updateUserBalance := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1",
		columnBalance, tableUsers, columnUserId)
	row := u.db.QueryRow(context.Background(), updateUserBalance, ub.UserID)
	err := row.Scan(&ub.Balance)
	if err != nil {
		return ub, err
	}
	return ub, nil
}

// TransferFunds - method for transferring funds between users
func (u *UserRepo) TransferFunds(t models.Transfer) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	updateSenderBalance := fmt.Sprintf("UPDATE %s SET %s=%s-$1 WHERE %s=$2",
		tableUsers, columnBalance, columnBalance, columnUserId)
	result, err := tx.Exec(context.Background(), updateSenderBalance, t.Amount, t.SenderID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback(context.Background())
		return fmt.Errorf("user with id %d does not exist", t.SenderID)
	}
	tr := models.Transaction{
		UserID:  t.SenderID,
		Amount:  -t.Amount,
		Date:    time.Now().Truncate(time.Second),
		Message: fmt.Sprintf("outgoing transfer to the user %d", t.ReceiverID),
	}
	err = u.addTransaction(tr)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	updateReceiverBalance := fmt.Sprintf("UPDATE %s SET %s=%s+$1 WHERE %s=$2",
		tableUsers, columnBalance, columnBalance, columnUserId)
	result, err = tx.Exec(context.Background(), updateReceiverBalance, t.Amount, t.ReceiverID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected = result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback(context.Background())
		return fmt.Errorf("user with id %d does not exist", t.ReceiverID)
	}
	tr = models.Transaction{
		UserID:  t.ReceiverID,
		Amount:  t.Amount,
		Date:    time.Now().Truncate(time.Second),
		Message: fmt.Sprintf("incoming transfer from user %d", t.SenderID),
	}
	err = u.addTransaction(tr)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

// UnblockFunds - method of reserving money if it was not possible to apply the service
func (u *UserRepo) UnblockFunds(unblock models.Unblock) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		return err
	}
	getAmount := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s=$1",
		columnAmount, columnUserId, tableOrders, columnOrderId)
	row := tx.QueryRow(context.Background(), getAmount, unblock.OrderID)
	err = row.Scan(&unblock.Amount, &unblock.UserID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	updateUserBalance := fmt.Sprintf("UPDATE %s SET %s=%s+$1 WHERE %s=$2",
		tableUsers, columnBalance, columnBalance, columnUserId)
	result, err := tx.Exec(context.Background(), updateUserBalance, unblock.Amount, unblock.UserID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		tx.Rollback(context.Background())
		return fmt.Errorf("user with id %d does not exist", unblock.UserID)
	}

	updateOrderStatus := fmt.Sprintf("UPDATE %s SET %s=$1,%s=$2 WHERE %s=$3",
		tableOrders, columnAmount, columnBlock, columnOrderId)
	result, err = tx.Exec(context.Background(), updateOrderStatus, 0, false, unblock.OrderID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	t := models.Transaction{
		UserID:  unblock.UserID,
		Amount:  unblock.Amount,
		Date:    time.Now().Truncate(time.Second),
		Message: "cancellation of service payment",
	}
	err = u.addTransaction(t)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

// addTransaction - adds information about the new transaction to the database
func (u *UserRepo) addTransaction(t models.Transaction) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	insertTx := fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s) VALUES ($1, $2, $3, $4)",
		tableTransactions, columnUserId, columnAmount, columnDate, columnMessage)
	result, err := tx.Exec(context.Background(), insertTx, t.UserID, t.Amount, t.Date, t.Message)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		tx.Rollback(context.Background())
		return errInsertRow

	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

// addUser - method for inserting a new user with replenished balance
func (u *UserRepo) addUser(ac models.AccrualCash) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	insertUser := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2)",
		tableUsers, columnUserId, columnBalance)
	result, err := tx.Exec(context.Background(), insertUser, ac.UserID, ac.Amount)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		tx.Rollback(context.Background())
		return errInsertRow
	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}
