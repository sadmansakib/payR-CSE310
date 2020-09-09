package services

import (
	"errors"
	"payR/database"
	"payR/models"
)

func getAllCustomers() []models.Customer {
	return database.Allcustomers
}

func getCustomerByID(id int) (models.Customer, error) {
	var foundCustomer models.Customer
	for _, customer := range database.Allcustomers {
		if id == customer.ID {
			foundCustomer = customer
			return foundCustomer, nil
		}
	}
	return foundCustomer, errors.New("Customer not found")
}
