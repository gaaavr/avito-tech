package models

import "time"

// AccrualCash - structure for a request to credit funds to a user's balance
type AccrualCash struct {
	ID      int     `json:"id" validate:"gte=1"`
	Amount  float64 `json:"amount" validate:"gt=0"`
	Message string
}

// UserBalance - structure to get the user's balance
type UserBalance struct {
	ID      int     `json:"id" validate:"gte=1"`
	Balance float64 `json:"balance"`
}

// Transaction - structure for transaction
type Transaction struct {
	UserID  int       `json:"id"`
	Amount  float64   `json:"amount"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}

// Order - structure for order
type Order struct {
	OrderID   int       `json:"order_id" validate:"gte=1"`
	UserID    int       `json:"id" validate:"gte=1"`
	ServiceID int       `json:"service_id" validate:"gte=1"`
	Amount    float64   `json:"amount" validate:"gt=0"`
	Date      time.Time `json:"date"`
	Block     bool
}
