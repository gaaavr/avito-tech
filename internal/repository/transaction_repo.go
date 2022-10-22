package repository

import (
	"avito/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

// TransactionRepo - transaction object in the repository layer
type TransactionRepo struct {
	db *pgxpool.Pool
}

// NewTransactionRepo - constructor function for TransactionRepo
func NewTransactionRepo(db *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{db: db}
}

// GetUserTransactions - method to get list of user's transactions
func (tx *TransactionRepo) GetUserTransactions(t models.TransactionListRequest) ([]models.TransactionList, error) {
	getTransactionsList := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1",
		tableTransactions, columnUserId)
	if strings.Contains(t.OrderBy, columnAmount) || strings.Contains(t.OrderBy, columnDate) {
		getTransactionsList += fmt.Sprintf(" ORDER BY %s", t.OrderBy)
	}
	if t.Limit != 0 {
		getTransactionsList += fmt.Sprintf(" LIMIT %d", t.Limit)
	}
	getTransactionsList += fmt.Sprintf(" OFFSET %d", t.Offset)
	rows, err := tx.db.Query(context.Background(), getTransactionsList, t.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tl := make([]models.TransactionList, 0, 1)
	for rows.Next() {
		tr := models.TransactionList{}
		err = rows.Scan(&tr.TransactionID, &tr.UserID, &tr.Amount, &tr.Date, &tr.Message)
		if err != nil {
			return nil, err
		}
		tl = append(tl, tr)
	}
	return tl, nil
}
