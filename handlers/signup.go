package handlers

import (
	"net/http"
	"payR/models"
	"payR/services"

	"github.com/gin-gonic/gin"
)

func SignupNewCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		customer := models.Customer{}
		c.ShouldBind(&customer)

		services.SignupCustomer(customer)

		c.Status(http.StatusCreated)
	}
}
