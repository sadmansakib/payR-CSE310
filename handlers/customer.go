package handlers

import (
	"database/sql"
	"net/http"
	"payR/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCustomers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"customers": services.GetAllCustomers(),
		})
	}
}

func GetCustomerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.MustGet("customer_id").(string) // Get the user id put in by auth middleware by decoding token
		client := c.MustGet("client").(*sql.DB)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
			return
		}
		customer, err := services.GetCustomerByID(id, client)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"customer": customer,
		})
	}
}
