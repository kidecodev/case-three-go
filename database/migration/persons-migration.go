package migration

import (
	"fmt"
	"frmgol/models"

	"gorm.io/gorm"
)

//PersonMigrate function
func PersonMigrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Persons{}) {
		db.Migrator().CreateTable(&models.Persons{})
		fmt.Println("Success Create Table Persons")
	}
}
