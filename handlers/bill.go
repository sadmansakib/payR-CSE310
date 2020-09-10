package handlers

import (
	"net/http"
	"payR/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBillsByCustomerID() gin.HandlerFunc {
	return func(c *gin.Context) {
		cstmrID := c.MustGet("customer_id").(string)
		id, err := strconv.Atoi(cstmrID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
			return
		}
		bill, err := services.GetBillsByCustomerID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No bills were submitted"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"bill": bill,
		})
	}
}
