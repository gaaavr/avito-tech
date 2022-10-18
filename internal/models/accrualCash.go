package models

// AccrualCash - structure for a request to credit funds to a user's balance
type AccrualCash struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}
