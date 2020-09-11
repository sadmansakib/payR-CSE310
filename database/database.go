package database

import "payR/models"

var customer1 = models.Customer{
	ID:             1,
	FName:          "customer",
	LName:          "one",
	Email:          "abc@example.com",
	HashedPassword: "asd",
	Mobile:         "0123456",
}

var customer2 = models.Customer{
	ID:             2,
	FName:          "customer",
	LName:          "two",
	Email:          "ksk@example.com",
	HashedPassword: "asd",
	Mobile:         "0123458",
}

var billCustomer1 = models.Bill{
	ID:             3,
	CustomerID:     1,
	Mobile:         "0123456",
	BillType:       "Postpaid",
	EquipmentCount: 1,
	Amount:         550.00,
	AccountID:      1,
	PaymentMethod:  "credit Card",
	SubmitDate:     "02/10/2020",
}

var billCustomer2 = models.Bill{
	ID:             25,
	CustomerID:     2,
	Mobile:         "0123458",
	BillType:       "Postpaid",
	EquipmentCount: 1,
	Amount:         550.00,
	AccountID:      1,
	PaymentMethod:  "bkash",
	SubmitDate:     "02/10/2020",
}

var accountState = models.Account{
	ID:          1,
	TotalAmount: 10500.00,
}

var Allcustomers = []models.Customer{
	customer1, customer2,
}

var AllBills = []models.Bill{
	billCustomer1, billCustomer2,
}
