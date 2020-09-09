package services

import (
	"log"
	"payR/database"
)

func SignupCustomer(
	firstName string,
	lastName string,
	email string,
	password string,
) {
	var db = database.ConnectDB()

	sqlQuery := `INSERT INTO customer (fname, lname, email, pass)
	VALUES ($1, $2, $3, $4)`

	_, error := db.Exec(sqlQuery, firstName, lastName, email, password)

	if error != nil {
		log.Fatal(error)
		panic(error)
	}
}
