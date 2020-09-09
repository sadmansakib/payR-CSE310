package services

import (
	"payR/database"
	"payR/models"
)

func GetAllBills() []models.Bill {
	return database.AllBills
}

func GetBillsByCustomerID(id int) models.Bill {
	var billofCustomer models.Bill
	for _, bill := range database.AllBills {
		if id == bill.CustomerID {
			billofCustomer = bill
		}
	}
	return billofCustomer
}
