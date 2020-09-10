package middleware

import (
	"fmt"
	"net/http"
	"payR/models"
	"strconv"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

var Hs = jwt.NewHS256([]byte("my_secret_key"))

func CheckAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AUTH_CHECK: called")
		// token is usually put in the header, not the json body
		token := []byte(c.GetHeader("token")) // the jwt library needs token to be in byte format, not string
		fmt.Printf("\ntoken: %s\n\n", string(token))
		if len(token) > 0 {
			//--------------------------------IF TOKEN PROVIDED
			var pl models.CustomPayload
			_, err := jwt.Verify(token, Hs, &pl)
			if err == nil {
				// if jwt verification is successful
				CustomerID, err := strconv.Atoi(pl.Subject)
				if err != nil {
					fmt.Println(CustomerID)
				}
				//
				// Query here like SELECT * FROM CUSTOMER WHERE Id = CustomerID
				//
				if false {
					// if no result (replace false with err != nil or whatever)
					fmt.Println("Invalid User")
					c.JSON(http.StatusInternalServerError, gin.H{"error": "User doesn't exist"})
					c.Abort() // Abort and return will cause the request to be dropped before it even reaches a handler
					return
				} else {
					fmt.Printf("\nvalid customer: id is %s\n\n", pl.Subject)

					c.Set("customer_id", pl.Subject)
					c.Next() // c.Next() will directly let the request go to the handler
					// Not mandatory, If the middleware doesn't abort, it will call c.Next() automatically
				}
			} else {
				c.JSON(http.StatusForbidden, gin.H{
					"error": err,
				})
				c.Abort()
				return
			}
		} else {
			//--------------------------------IF TOKEN NOT PROVIDED
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "You must provide an authorization token",
			})
			c.Abort()
			return
		}
	}
}
