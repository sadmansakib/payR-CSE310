package models

type Account struct {
	ID          int     `json:"account_id"`
	TotalAmount float64 `json:"total_amount"`
}
