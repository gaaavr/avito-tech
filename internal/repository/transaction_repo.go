package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// TransactionRepo - transaction object in the repository layer
type TransactionRepo struct {
	db *pgxpool.Pool
}

// NewTransactionRepo - constructor function for UserRepo
func NewTransactionRepo(db *pgxpool.Pool) *TransactionRepo {
	return &TransactionRepo{db: db}
}
