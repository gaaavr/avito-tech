package service

import (
	"avito/internal/models"
	"avito/internal/repository"
	"errors"
)

var errEmptyList = errors.New("user has no transactions")

// TransactionService - transaction object in the service layer
type TransactionService struct {
	repo repository.Transaction
}

// NewTransactionService - constructor function for TransactionService
func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

// GetUserTransactions - method to get list of user's transactions
func (t *TransactionService) GetUserTransactions(tr models.TransactionListRequest) (tl []models.TransactionList, code int, err error) {
	if tr.UserID < 1 {
		return nil, 400, errUser
	}
	tl, err = t.repo.GetUserTransactions(tr)
	if err != nil {
		return nil, 500, err
	}
	if len(tl) == 0 {
		return nil, 400, errEmptyList
	}
	return tl, 200, nil
}
