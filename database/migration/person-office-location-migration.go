package migration

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
)

func personOfficeLocationMigrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.OfficePersonLocations{}) {
		db.Migrator().CreateTable(&models.OfficePersonLocations{})
		fmt.Println("Success Create Table Office Person Locations")
	}
}