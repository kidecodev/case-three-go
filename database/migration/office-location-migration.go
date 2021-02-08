package migration

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
)

func officeLocationMigrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.OfficeLocations{}) {
		db.Migrator().CreateTable(&models.OfficeLocations{})
		fmt.Println("Success Create Table Office Locations")
	}
}