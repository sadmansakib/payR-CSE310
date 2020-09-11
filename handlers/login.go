package handlers

import (
	"fmt"
	"log"
	"net/http"
	"payR/database"
	"payR/middleware"
	"payR/models"
	"strconv"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

type EmailAndPass struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("LOGIN: called")
		emailAndPass := EmailAndPass{}

		err := c.ShouldBindJSON(&emailAndPass)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(emailAndPass)
		email, password := emailAndPass.Email, emailAndPass.Password

		if len(email) > 0 && len(password) > 0 {
			//--------------------------------IF ALL FIELDS PROVIDED
			//
			// CHECK Db for user with email and password
			//
			if false {
				//--------------------------------IF USER DOESN'T EXIST
				fmt.Println("LOGIN: User Doesn't Exist - Aborting Login")
				c.JSON(http.StatusNotFound, gin.H{"error": "User doesn't exist"})
			} else {
				//--------------------------------IF USER EXISTS
				fmt.Println("LOGIN: User Exists - Attempting Login")
				customer := models.Customer{}
				//
				// Bind the DB data to customer here
				//
				customer = database.Allcustomers[0] // using the dummy customer for now
				//
				if true {
					//--------------------------------IF USER OBJECT CREATED SUCCESSFULLY
					fmt.Println("LOGIN: User decoded")
					if customer.Password == password {
						//--------------------------------IF PASSWORD MATCHES
						fmt.Println("LOGIN: Password matched")
						now := time.Now()
						pl := models.CustomPayload{
							Payload: jwt.Payload{
								Issuer:         "gbrlsnchs",
								Subject:        strconv.Itoa(customer.ID), //Only this field is relevant to us, other are fancy stuff I have no idea about
								Audience:       jwt.Audience{"https://golang.org", "https://jwt.io"},
								ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
								NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
								IssuedAt:       jwt.NumericDate(now),
								JWTID:          "foobar",
							},
						}

						token, err := jwt.Sign(pl, middleware.Hs)
						if err != nil {
							//--------------------------------IF JWT SIGNING FAILS
							c.JSON(http.StatusInternalServerError, gin.H{
								"error": err,
							})
						} else {
							//--------------------------------IF JWT SIGNING SUCCEEDES
							fmt.Println(gin.H{
								"token": string(token),
							})
							c.JSON(http.StatusOK, gin.H{
								"token": string(token),
							})
						}

					} else {
						//--------------------------------IF PASSWORD DEOSN'T MATCH
						fmt.Println("LOGIN: Password Mismatch")
						c.JSON(http.StatusUnauthorized, gin.H{
							"error": "password mismatch",
						})
					}
				} else {
					//--------------------------------IF USER OBJECT DECODE FAILS
					fmt.Println("LOGIN: User Decode Failed")
					log.Fatal("usererr")
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "couldn't decode user",
					})
				}
			}
		} else {
			//--------------------------------IF ALL FIELDS NOT PROVIDED
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "You must provide email and password",
			})
		}
	}
}
