package repository

import (
	"avito/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// User - Interface describing the user entity
type User interface {
	AccrualFunds(ac models.AccrualCash) error
	GetBalance(ub *models.UserBalance) (*models.UserBalance, error)
	BlockFunds(order models.Order) error
}

// Transaction - interface describing the transaction object
type Transaction interface {
}

// Order - interface describing the Order object
type Order interface {
	ChargeFunds(order models.Order) error
}

// Repository - object responsible for the work of logic with the database
type Repository struct {
	User
	Transaction
	Order
}

// NewRepository - constructor function for Repository
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User:        NewUserRepo(db),
		Transaction: NewTransactionRepo(db),
		Order:       NewOrdersRepo(db),
	}
}
