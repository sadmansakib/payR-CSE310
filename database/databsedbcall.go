package database

func getAllCustomers() []CustomerMock {
	return allcustomers
}

func getCustomerByID(id int) CustomerMock {
	var foundCustomer CustomerMock
	for _, customer := range allcustomers {
		if id == customer.id {
			foundCustomer = customer
		}
	}
	return foundCustomer
}

func getAllBills() []Bill {
	return allBills
}

func getBillsByCustomerID(id int) Bill {
	var billofCustomer Bill
	for _, bill := range allBills {
		if id == bill.customerID {
			billofCustomer = bill
		}
	}
	return billofCustomer
}
