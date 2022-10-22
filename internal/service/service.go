// Package service is responsible for business logic, contains all the abstractions needed on this layer
package service

import (
	"avito/internal/models"
	"avito/internal/repository"
)

// User - Interface describing the user entity
type User interface {
	AccrualFunds(ac models.AccrualFunds) (code int, err error)
	GetBalance(ub *models.UserBalance) (code int, err error)
	BlockFunds(order models.Order) (code int, err error)
	TransferFunds(t models.Transfer) (code int, err error)
	UnblockFunds(unblock models.Unblock) (code int, err error)
}

// Order - Interface describing the order entity
type Order interface {
	ChargeFunds(order models.Order) (code int, err error)
	GetReport(report models.Report) (data []byte, code int, err error)
}

// Transaction - Interface describing the transaction entity
type Transaction interface {
	GetUserTransactions(tr models.TransactionListRequest) (tl []models.TransactionList, code int, err error)
}

// Service - object responsible for the operation of the internal logic
type Service struct {
	User
	Order
	Transaction
}

// NewService - constructor function for Service
func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:        NewUserService(repository.User),
		Order:       NewOrderService(repository.Order),
		Transaction: NewTransactionService(repository.Transaction),
	}
}
