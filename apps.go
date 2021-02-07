package main

import (
	"frmgol/config"
	"frmgol/database/migration"
	"frmgol/database/seed"
	"frmgol/web"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	migration.Migrations()
	seed.Seeder()
	router := web.Route(gin.Default())
	router.Run(":8080")
}
