package service

import (
	"avito/internal/models"
	"avito/internal/repository"
	"errors"
	"fmt"
	"time"
)

var (
	errAmount  = errors.New("the amount of funds must be greater than 0")
	errUser    = errors.New("user id must not be less than 1")
	errService = errors.New("service id must not be less than 1")
	errOrder   = errors.New("order id must not be less than 1")
	errNoRows  = "no rows in result set"
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
		return 400, errAmount
	}
	if ac.ID < 1 {
		return 400, errUser
	}
	ac.Message = "replenishment of the balance"
	err = u.repo.AccrualFunds(ac)
	if err != nil {
		return 500, fmt.Errorf("database error: %s", err.Error())
	}
	return 200, nil
}

// BlockFunds - method of reserving funds from the main balance in a separate account
func (u *UserService) BlockFunds(order models.Order) (code int, err error) {
	if order.Amount <= 0 {
		return 400, errAmount
	}
	if order.UserID < 1 {
		return 400, errUser
	}
	if order.ServiceID < 1 {
		return 400, errService
	}
	if order.OrderID < 1 {
		return 400, errOrder
	}
	order.Date = time.Now().Truncate(time.Second)
	order.Block = true
	err = u.repo.BlockFunds(order)
	if err != nil {
		if err.Error() == errNoRows {
			return 400, fmt.Errorf("user with id %d does not exist", order.UserID)
		}
		return 500, fmt.Errorf("database error: %s", err.Error())
	}
	return 200, nil
}

// GetBalance - method to get user balance
func (u *UserService) GetBalance(ub *models.UserBalance) (code int, err error) {
	if ub.ID < 1 {
		return 400, errUser
	}
	ub, err = u.repo.GetBalance(ub)
	if err != nil {
		if err.Error() == errNoRows {
			return 400, fmt.Errorf("user with id %d does not exist", ub.ID)
		}
		return 500, fmt.Errorf("database error: %s", err.Error())
	}

	return 200, nil
}
