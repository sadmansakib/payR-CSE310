package services

import (
	"database/sql"
	"payR/models"
)

func GetCustomerByID(id int, client *sql.DB) (models.Customer, error) {
	var foundCustomer models.Customer
	sqlQuery :=
		`
	SELECT 
		id,
		email,
		fName,
		lName,
		mobile
	FROM 
		Customer
	WHERE 
		id = $1;
	`

	row := client.QueryRow(sqlQuery, id)

	err := row.Scan(
		&foundCustomer.ID,
		&foundCustomer.Email,
		&foundCustomer.FName,
		&foundCustomer.LName,
		&foundCustomer.Mobile,
	)

	return foundCustomer, err
}
