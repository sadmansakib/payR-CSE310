package models

type Customer struct {
	ID       int    `json:"id,omitempty"`
	FName    string `json:"first_name"`
	LName    string `json:"last_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type Bill struct {
	ID             int     `json:"bill_id"`
	CustomerID     int     `json:"id,omitempty"`
	Mobile         string  `json:"mobile"`
	BillType       string  `json:"bill_type"`
	EquipmentCount int     `json:"equipment_count"`
	Amount         float64 `json:"amount"`
	AccountID      int     `json:"account_id"`
	PaymentMethod  string  `json:"payment_method"`
	SubmitDate     string  `json:"submition_data"`
}

type Account struct {
	ID          int     `json:"account_id"`
	TotalAmount float64 `json:"total_amount"`
}
