package migration

import (
	"fmt"
	"frmgol/models"

	"gorm.io/gorm"
)

func provinceMigrate(db *gorm.DB) {

	if !db.Migrator().HasTable(&models.Provinces{}) {
		db.Migrator().CreateTable(&models.Provinces{})
		fmt.Println("Success Create Table Provinces")
	}
}
