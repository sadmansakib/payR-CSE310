package handlers

import (
	"database/sql"
	"net/http"
	"payR/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBillsByCustomerID() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := c.MustGet("client").(*sql.DB)
		cstmrID := c.MustGet("customer_id").(string)
		id, err := strconv.Atoi(cstmrID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
			return
		}
		bill, err := services.GetBillsByCustomerID(id, client)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No bills were submitted"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"bill": bill,
		})
	}
}
