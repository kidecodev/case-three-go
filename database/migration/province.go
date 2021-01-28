package migration

import (
	"fmt"
	"frmgol/models"

	"gorm.io/gorm"
)

//ProvinceMigrate function
func ProvinceMigrate(db *gorm.DB) {

	if db.Migrator().HasTable(&models.Provinces{}) {
		db.Migrator().CreateTable(&models.Provinces{})
		fmt.Println("Success Create Table Provinces")
	}
}
