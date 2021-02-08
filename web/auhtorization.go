package web

import (
	"fmt"
	"frmgol/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authorization(context *gin.Context) {
	var (
		token_from_auth string
		method string
		secret string
	)
	secret = utils.GetEnv("SECRET_TOKEN")
	token_from_auth = context.Request.Header.Get("Authorization")
	method = context.Request.Method
	
	token, err := jwt.Parse(token_from_auth, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Method false $e", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if token != nil && err == nil {
		paylod := token.Claims.(jwt.MapClaims)
		if paylod["role"] == "guest" && method != "GET" {

		} else if paylod["role"] == "entry" && method == "DELETE" {

		} else {
			context.Next()
		}
	} else {
		result := gin.H{
			"message": "Token false",
			"error": http.StatusUnauthorized,
		}

		context.Abort()
		context.JSON(http.StatusUnauthorized, result)
	}
}
