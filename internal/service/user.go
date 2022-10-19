package service

import (
	"avito/internal/models"
	"avito/internal/repository"
	"fmt"
)

// UserService - user object in the service layer
type UserService struct {
	repo repository.User
}

// NewUserService - constructor function for UserService
func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

// AccrualFunds - method of accruing cash to the balance
func (u *UserService) AccrualFunds(ac models.AccrualCash) (code int, err error) {
	if ac.Amount <= 0 {
		return 400, fmt.Errorf("the amount of funds to be credited must be greater than 0")
	}
	if ac.ID <= 0 {
		return 400, fmt.Errorf("user id must be greater than 0")
	}
	if ac.Message == "" {
		ac.Message = "replenishment of the balance"
	}
	err = u.repo.AccrualFunds(ac)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
