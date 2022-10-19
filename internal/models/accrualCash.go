package models

// AccrualCash - structure for a request to credit funds to a user's balance
type AccrualCash struct {
	ID      int     `json:"id" validate:"required,gte=1"`
	Amount  float64 `json:"amount" validate:"required,gte=1"`
	Message string  `json:"message"`
}
