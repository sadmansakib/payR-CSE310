package models

import (
	"time"
)

type Bill struct {
	ID             int     `json:"bill_id"`
	CustomerID     int     `json:"id,omitempty"`
	Mobile         string  `json:"mobile"`
	BillType       string  `json:"bill_type"`
	EquipmentCount int     `json:"equipment_count"`
	Amount         float64 `json:"amount"`
	AccountID      int     `json:"account_id"`
	PaymentMethod  string  `json:"payment_method"`
	SubmitDate     time.Time
}
