package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"payR/models"
	"payR/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CustomerBinding struct {
	FName    string `json:"first_name" form:"first_name"`
	LName    string `json:"last_name" form:"last_name"`
	Email    string `json:"email" form:"email"`
	Mobile   string `json:"mobile" form:"mobile"`
	Password string `json:"password" form:"password"`
}

func SignupNewCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerBinding := CustomerBinding{}
		client := c.MustGet("client").(*sql.DB)

		c.ShouldBind(&customerBinding)

		passwrd, error := bcrypt.GenerateFromPassword([]byte(customerBinding.Password), bcrypt.DefaultCost)

		if error != nil {
			fmt.Println(error)
		}

		customerBinding.Password = string(passwrd)

		customer := models.Customer{
			FName:          customerBinding.FName,
			LName:          customerBinding.LName,
			Email:          customerBinding.Email,
			Mobile:         customerBinding.Mobile,
			HashedPassword: customerBinding.Password,
		}

		services.SignupCustomer(customer, client)

		c.Status(http.StatusCreated)
	}
}
