package services

import (
	"errors"
	"payR/database"
	"payR/models"
)

func GetBillsByCustomerID(id int) (models.Bill, error) {
	var billofCustomer models.Bill
	for _, bill := range database.AllBills {
		if id == bill.CustomerID {
			billofCustomer = bill
			return billofCustomer, nil
		}
	}
	return billofCustomer, errors.New("No bills were found")
}
