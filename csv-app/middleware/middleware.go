package middleware

import (
	dbconnect "csv-app/db-connect"

	"github.com/gin-gonic/gin"
)

func InitDatabase() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := dbconnect.ConnectToDatabase()

		ctx.Set("db", db)
	}
}
