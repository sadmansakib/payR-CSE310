package main

import (
	"payR/database"
	"payR/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var db = database.ConnectDB()

	router := gin.Default()

	router.GET("/customers", handlers.GetAllCustomers())
	router.GET("/customers/:id", handlers.GetCustomerById())

	router.GET("/customers/:id/bills", handlers.GetBillsByCustomerID())

	router.Run()

	println(db.Stats)
}
