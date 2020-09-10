package services

import (
	"fmt"
	"log"
	"payR/database"
	"payR/models"
)

func SignupCustomer(
	customer models.Customer,
) {
	fmt.Println(customer)
	var db = database.ConnectDB()
	sqlQuery := `INSERT INTO customer (fname, lname, email, pass)
	VALUES ($1, $2, $3, $4)`

	_, error := db.Exec(sqlQuery, customer.FName, customer.LName, customer.Email, customer.Password)

	if error != nil {
		log.Fatal(error)
		panic(error)
	}
}
