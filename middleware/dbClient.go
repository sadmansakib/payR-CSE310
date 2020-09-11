package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func ProvideDBInstance(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("client", db)
	}
}
