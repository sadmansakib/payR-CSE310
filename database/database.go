package database

type CustomerMock struct {
	id       int    `json:"id,omitempty"`
	fName    string `json:"first_name"`
	lName    string `json:"last_name"`
	email    string `json:"email"`
	mobile   string `json:"mobile"`
	password string `json:"password"`
}

type Customer struct {
	fName    string `json:"first_name"`
	lName    string `json:"last_name"`
	email    string `json:"email"`
	mobile   string `json:"mobile"`
	password string `json:"password"`
}

type Bill struct {
	id             int     `json:"bill_id"`
	customerID     int     `json:"id,omitempty"`
	mobile         string  `json:"mobile"`
	billType       string  `json:"bill_type"`
	equipmentCount int     `json:"equipment_count"`
	amount         float64 `json:"amount"`
	accountID      int     `json:"account_id"`
	paymentMethod  string  `json:"payment_method"`
	submitDate     string  `json:"submition_data"`
}

type Account struct {
	id          int     `json:"account_id"`
	totalAmount float64 `json:"total_amount"`
}
