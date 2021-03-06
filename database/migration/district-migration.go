package migration

import (
	"fmt"
	"frmgol/models"

	"gorm.io/gorm"
)

func districtMigrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Districts{}) {
		db.Migrator().CreateTable(&models.Districts{})
		fmt.Println("Success Create Table Districts")
	}
}
