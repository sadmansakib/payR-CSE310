package handlers

import (
	"database/sql"
	"net/http"
	"payR/models"
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

		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusNotFound, gin.H{"error": "No bills where submitted"})
			return
		case nil:
			c.JSON(http.StatusOK, gin.H{"bills": bill})
			return
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "No bills where submitted"})
		}
	}
}

type BillBinding struct {
	CustomerID     int     `json:"id,omitempty"`
	Mobile         string  `json:"mobile"`
	BillType       string  `json:"bill_type"`
	EquipmentCount int     `json:"equipment_count"`
	Amount         float64 `json:"amount"`
	AccountID      int     `json:"account_id"`
	PaymentMethod  string  `json:"payment_method"`
	SubmitDate     string  `json:"submition_data"`
}

func SubmitBill() gin.HandlerFunc {
	return func(c *gin.Context) {
		billBinding := BillBinding{}
		client := c.MustGet("client").(*sql.DB)
		cstmrID := c.MustGet("customer_id").(string)
		id, err := strconv.Atoi(cstmrID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
			return
		}

		c.ShouldBindJSON(&billBinding)

		bill := models.Bill{
			CustomerID:     id,
			Mobile:         billBinding.Mobile,
			BillType:       billBinding.BillType,
			EquipmentCount: billBinding.EquipmentCount,
			Amount:         billBinding.Amount,
			AccountID:      1,
			PaymentMethod:  billBinding.PaymentMethod,
			SubmitDate:     billBinding.SubmitDate,
		}

		services.SubmitBill(bill, client)

		c.JSON(http.StatusOK, gin.H{"success": "Bill payment successful"})

	}
}
