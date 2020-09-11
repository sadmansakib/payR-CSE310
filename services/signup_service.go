package services

import (
	"database/sql"
	"log"
	"payR/models"
)

func SignupCustomer(
	customer models.Customer,
	dbInstance *sql.DB,
) {

	sqlQuery := `INSERT INTO customer (fname, lname, email, mobile, pass)
	VALUES ($1, $2, $3, $4, $5)`

	_, error := dbInstance.Exec(sqlQuery,
		customer.FName,
		customer.LName,
		customer.Email,
		customer.Mobile,
		customer.HashedPassword)

	if error != nil {
		log.Fatal(error)
		panic(error)
	}
}
