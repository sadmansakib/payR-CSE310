package handlers

import (
	"net/http"
	"payR/services"

	"github.com/gin-gonic/gin"
)

func SignupNewCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		fName := c.Query("first_name")
		lName := c.Query("last_name")
		mail := c.Query("email")
		pass := c.Query("password")

		services.SignupCustomer(fName, lName, mail, pass)

		c.String(http.StatusCreated, "SUCCESS")
	}
}
