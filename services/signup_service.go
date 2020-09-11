package services

import (
	"log"
	db "payR/database"
	"payR/models"
)

func SignupCustomer(
	customer models.Customer,
) {

	sqlQuery := `INSERT INTO customer (fname, lname, email, mobile, pass)
	VALUES ($1, $2, $3, $4, $5)`

	_, error := db.DBInstance.Exec(sqlQuery,
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
