package main

import (
	"payR/database"
	"payR/handlers"
	"payR/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.ProvideDBInstance(database.DBInstance))
	router.Use(middleware.CorsMiddleware())

	//--------------------------Authentication Required Routes
	authorized := router.Group("/")

	// All the routes under `authorized` use the CheckAuthorization middleware
	// The middleware can
	// - let the request get to the handler `by calling c.Next()`
	// - modify the request (example: decode the token and add the decoded customer_id to the request body) `by calling c.Set()`
	// - completely abort the request, not even let it reach the handler `by calling c.Abort() and return`
	// CHECK THE CheckAuthorization function for details
	authorized.Use(middleware.CheckAuthorization())
	{
		authorized.GET("/customers/customer", handlers.GetCustomerById())
		authorized.GET("/customers/customer/bills", handlers.GetBillsByCustomerID())
		authorized.POST("/customers/customer/bills", handlers.SubmitBill())
	}

	router.POST("/signup", handlers.SignupNewCustomer())
	router.POST("/login", handlers.Login())

	router.Run() //you can do `router.Run(":1234")` to change port

}
