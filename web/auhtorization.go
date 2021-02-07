package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authorization(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
			"error":   "token salah",
		})

		context.Abort()
	}
}
