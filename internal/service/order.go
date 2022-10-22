package service

import (
	"avito/internal/models"
	"avito/internal/repository"
	"errors"
	"fmt"
	"strings"
)

var (
	errNoOrder = "this order does not exist"
	errYear    = errors.New("year must not be less than 2007")
	errMonth   = errors.New("month must be between 1 and 12")
	errNoData  = errors.New("users did not use the services at this time")
)

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

// GetReport - method for providing a summary report for accounting
func (o *OrderService) GetReport(report models.Report) (data []byte, code int, err error) {
	if report.Year < 2007 {
		return nil, 400, errYear
	}
	if report.Month < 1 || report.Month > 12 {
		return nil, 400, errMonth
	}
	report.Data = make(map[int]float64)
	err = o.repo.GetReport(&report)
	if err != nil {
		return nil, 500, err
	}
	if len(report.Data) == 0 {
		return nil, 400, errNoData
	}
	var sb strings.Builder
	for serviceID, amount := range report.Data {
		str := fmt.Sprintf("%d;%0.2f\n", serviceID, amount)
		_, err = sb.WriteString(str)
		if err != nil {
			return nil, 500, err
		}
	}
	data = []byte(strings.TrimSuffix(sb.String(), "\n"))
	return data, 200, nil
}
