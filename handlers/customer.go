package handlers

import (
	"database/sql"
	"net/http"
	"payR/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
			return
		case nil:
			c.JSON(http.StatusOK, gin.H{"customer": customer})
			return
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		}

	}
}
