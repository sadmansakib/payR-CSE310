package main

import (
	"payR/handlers"
	"payR/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/customers", handlers.GetAllCustomers())

	//--------------------------Authentication Required Routes
	authorized := router.Group("/")
	authorized.Use(middleware.CheckAuthorization())
	{
		authorized.GET("/customers/customer", handlers.GetCustomerById())

		authorized.GET("/customers/customer/bills", handlers.GetBillsByCustomerID())
	}

	router.POST("/signup", handlers.SignupNewCustomer())
	router.POST("/login", handlers.Login())

	router.Run()

}
