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
func (u *UserService) AccrualFunds(ac models.AccrualFunds) (code int, err error) {
	if ac.Amount <= 0 {
		return 400, errAmount
	}
	if ac.UserID < 1 {
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
	return 201, nil
}

// UnblockFunds - method of reserving money if it was not possible to apply the service
func (u *UserService) UnblockFunds(unblock models.Unblock) (code int, err error) {
	if unblock.OrderID < 1 {
		return 400, errOrder
	}
	err = u.repo.UnblockFunds(unblock)
	if err != nil {
		if err.Error() == errNoRows {
			return 400, fmt.Errorf("it is not possible to unlock funds for this order")
		}
		return 500, fmt.Errorf("database error: %s", err.Error())
	}
	return 200, nil
}

// GetBalance - method to get user balance
func (u *UserService) GetBalance(ub *models.UserBalance) (code int, err error) {
	if ub.UserID < 1 {
		return 400, errUser
	}
	ub, err = u.repo.GetBalance(ub)
	if err != nil {
		if err.Error() == errNoRows {
			return 400, fmt.Errorf("user with id %d does not exist", ub.UserID)
		}
		return 500, fmt.Errorf("database error: %s", err.Error())
	}

	return 200, nil
}

// TransferFunds - method for transferring funds between users
func (u *UserService) TransferFunds(t models.Transfer) (code int, err error) {
	if t.SenderID < 1 || t.ReceiverID < 1 {
		return 400, errUser
	}
	if t.Amount <= 0 {
		return 400, errAmount
	}
	err = u.repo.TransferFunds(t)
	if err != nil {
		return 500, fmt.Errorf("database error: %s", err.Error())
	}
	return 200, nil
}
