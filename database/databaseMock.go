package database

var customer1 = CustomerMock{
	id:       1,
	fName:    "abc",
	lName:    "edf",
	email:    "abc@example.com",
	password: "asd",
	mobile:   "0123456",
}

var customer2 = CustomerMock{
	id:       2,
	fName:    "edf",
	lName:    "ghi",
	email:    "ksk@example.com",
	password: "asd",
	mobile:   "0123458",
}

var billCustomer1 = Bill{
	id:             3,
	customerID:     1,
	mobile:         "0123456",
	billType:       "Postpaid",
	equipmentCount: 1,
	amount:         550.00,
	accountID:      1,
	paymentMethod:  "credit Card",
	submitDate:     "02/10/2020",
}

var billCustomer2 = Bill{
	id:             25,
	customerID:     2,
	mobile:         "0123458",
	billType:       "Postpaid",
	equipmentCount: 1,
	amount:         550.00,
	accountID:      1,
	paymentMethod:  "bkash",
	submitDate:     "02/10/2020",
}

var accountState = Account{
	id:          1,
	totalAmount: 10500.00,
}

var allcustomers = []CustomerMock{
	customer1, customer2,
}

var allBills = []Bill{
	billCustomer1, billCustomer2,
}
