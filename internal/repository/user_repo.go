package repository

import (
	"avito/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableUsers    = "users"
	columnBalance = "balance"
	columnUserId  = "user_id"
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
	result, err := tx.Exec(context.Background(), updateUserBalance, ac.Amount, ac.ID)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		err = u.insertUser(ac)
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	return nil
}

// insertUser - method for inserting a new user with replenished balance
func (u *UserRepo) insertUser(ac models.AccrualCash) error {
	tx, err := u.db.Begin(context.Background())
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	insertUser := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2)",
		tableUsers, columnUserId, columnBalance)
	result, err := tx.Exec(context.Background(), insertUser, ac.ID, ac.Amount)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
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
