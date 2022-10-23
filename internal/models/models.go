// Package models contains all the necessary objects for transferring data and interacting with the application
package models

import "time"

// AccrualFunds - structure for a request to credit funds to a user's balance
type AccrualFunds struct {
	UserID  int     `json:"user_id" validate:"gte=1"`
	Amount  float64 `json:"amount" validate:"gt=0"`
	Message string
}

// UserBalance - structure to get the user's balance
type UserBalance struct {
	UserID  int     `json:"user_id" validate:"gte=1"`
	Balance float64 `json:"balance"`
}

// Transaction - structure for transaction
type Transaction struct {
	UserID  int       `json:"user_id"`
	Amount  float64   `json:"amount"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}

// Order - an object for reserving funds when creating an order and debiting these funds
type Order struct {
	OrderID   int       `json:"order_id" validate:"gte=1"`
	UserID    int       `json:"user_id" validate:"gte=1"`
	ServiceID int       `json:"service_id" validate:"gte=1"`
	Amount    float64   `json:"amount" validate:"gt=0"`
	Date      time.Time `json:"date"`
	Block     bool
}

// Report - object for working with a report for accounting
type Report struct {
	Year  int `json:"year" validate:"required,gte=2007"`
	Month int `json:"month" validate:"required,min=1,max=12"`
	Data  map[int]float64
}

// Transfer - object for working with the transfer of funds between two users
type Transfer struct {
	SenderID   int     `json:"sender_id" validate:"gte=1"`
	ReceiverID int     `json:"receiver_id" validate:"gte=1"`
	Amount     float64 `json:"amount" validate:"gt=0"`
}

// TransactionList - object for working with the list of user transactions
type TransactionList struct {
	TransactionID int       `json:"transaction_id"`
	UserID        int       `json:"user_id"`
	Amount        float64   `json:"amount"`
	Date          time.Time `json:"date"`
	Message       string    `json:"message"`
}

// TransactionListRequest - structure for requesting a list of user transactions
type TransactionListRequest struct {
	UserID  int    `json:"user_id" validate:"gte=1"`
	OrderBy string `json:"order_by"`
	Limit   int    `json:"limit" validate:"gte=0"`
	Offset  int    `json:"offset" validate:"gte=0"`
}

// Unblock - structure for unlocking funds
type Unblock struct {
	OrderID int     `json:"order_id" validate:"gte=1"`
	UserID  int     `json:"user_id" validate:"gte=1"`
	Amount  float64 `json:"amount"`
}
