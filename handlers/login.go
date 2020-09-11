package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	db "payR/database"
	"payR/middleware"
	"payR/models"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

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

			sqlQuery := `SELECT id, pass FROM customer WHERE email = $1;`

			var fetchedCustomer models.Customer

			row := db.DBInstance.QueryRow(sqlQuery, email)

			err := row.Scan(&fetchedCustomer.ID, &fetchedCustomer.HashedPassword)

			fmt.Println(fetchedCustomer.ID)

			switch err {
			case sql.ErrNoRows:
				log.Fatal("LOGIN: User Doesn't Exist - Aborting Login")
				c.JSON(http.StatusNotFound, gin.H{"error": "User doesn't exist"})
				return
			case nil:
				errPass := bcrypt.CompareHashAndPassword([]byte(fetchedCustomer.HashedPassword), []byte(password))
				if errPass != nil && errPass == bcrypt.ErrMismatchedHashAndPassword {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect Password"})
				} else {
					fmt.Println("LOGIN: Password matched")
					now := time.Now()
					pl := models.CustomPayload{
						Payload: jwt.Payload{
							Issuer:         "gbrlsnchs",
							Subject:        strconv.Itoa(fetchedCustomer.ID), //Only this field is relevant to us, other are fancy stuff I have no idea about
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
				}
			default:
				panic(err)
			}
		} else {
			//--------------------------------IF ALL FIELDS NOT PROVIDED
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "You must provide email and password",
			})
		}
	}
}
