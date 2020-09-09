package main

import (
	"payR/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/customers", handlers.GetAllCustomers())
	router.GET("/customers/:id", handlers.GetCustomerById())

	router.GET("/customers/:id/bills", handlers.GetBillsByCustomerID())

	router.POST("/signup", handlers.SignupNewCustomer())

	router.Run()

}
