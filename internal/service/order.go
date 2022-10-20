package service

import (
	"avito/internal/models"
	"avito/internal/repository"
	"fmt"
)

var errNoOrder = "this order does not exist"

// OrderService - order object in the service layer
type OrderService struct {
	repo repository.Order
}

// NewOrderService - constructor function for OrderService
func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

// ChargeFunds - method for charging previously reserved funds
func (o *OrderService) ChargeFunds(order models.Order) (code int, err error) {
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
	order.Block = false
	err = o.repo.ChargeFunds(order)
	if err != nil {
		if err.Error() == errNoOrder {
			return 400, err
		}
		return 500, fmt.Errorf("database error: %s", err.Error())
	}
	return 200, nil
}
