package services

import (
	"database/sql"
	"errors"
	"log"
	"payR/models"
)

func GetCustomerByID(id int, client *sql.DB) (models.Customer, error) {
	var foundCustomer models.Customer
	sqlQuery := `SELECT * FROM customer WHERE id = $1;`

	row := client.QueryRow(sqlQuery, id)

	err := row.Scan(&foundCustomer.ID,
		&foundCustomer.FName,
		&foundCustomer.LName,
		&foundCustomer.Mobile,
		&foundCustomer.Email,
		&foundCustomer.HashedPassword)

	switch err {
	case sql.ErrNoRows:
		log.Fatal("LOGIN: User Doesn't Exist - Aborting Login")
		return foundCustomer, errors.New("Customer not found")
	case nil:
		return foundCustomer, nil
	default:
		panic(err)
	}
}
