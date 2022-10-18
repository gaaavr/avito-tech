package service

import (
	"avito/internal/models"
	"avito/internal/repository"
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

// AccrualCash - method of accruing cash to the balance
func (u *UserService) AccrualFunds(ac models.AccrualCash) (code int, err error) {
	u.repo.AccrualFunds(ac)
	return 200, nil
}
