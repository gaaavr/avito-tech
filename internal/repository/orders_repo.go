package repository

import (
	"avito/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableOrders     = "orders"
	columnBlock     = "block"
	columnOrderId   = "order_id"
	columnServiceId = "service_id"
)

// OrdersRepo - orders object in the repository layer
type OrdersRepo struct {
	db *pgxpool.Pool
}

// NewOrdersRepo - constructor function for OrdersRepo
func NewOrdersRepo(db *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{db: db}
}

// ChargeFunds - method for charging previously reserved funds
func (o *OrdersRepo) ChargeFunds(order models.Order) error {
	updateOrderStatus := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE %s=$2 AND %s=$3 AND %s=$4 AND %s=$5 AND %s=$6",
		tableOrders, columnBlock, columnOrderId, columnUserId, columnServiceId, columnAmount, columnBlock)
	result, err := o.db.Exec(context.Background(), updateOrderStatus, order.Block, order.OrderID, order.UserID,
		order.ServiceID, order.Amount, true)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("order does not exist")
	}
	return nil
}

// GetReport - method for providing a summary report for accounting
func (o *OrdersRepo) GetReport(report *models.Report) error {
	getReport := fmt.Sprintf("SELECT %s, sum(%s) FROM %s  WHERE DATE_PART('year',%s)=$1 AND DATE_PART('month',%s)=$2 AND %s<>$3 GROUP BY %s",
		columnServiceId, columnAmount, tableOrders, columnDate, columnDate, columnAmount, columnServiceId)
	rows, err := o.db.Query(context.Background(), getReport, report.Year, report.Month, 0)
	if err != nil {
		return err
	}
	defer rows.Close()
	var (
		serviceID int
		amount    float64
	)
	for rows.Next() {
		err = rows.Scan(&serviceID, &amount)
		if err != nil {
			return err
		}
		report.Data[serviceID] = amount
	}
	return nil
}
