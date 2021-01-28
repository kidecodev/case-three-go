package main

import (
	"frmgol/web"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	migration.Migrations()
	router := web.Route(gin.Default())
	router.Run(":8080")

}
