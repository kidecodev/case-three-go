package migration

import (
	"fmt"
	"frmgol/models"

	"gorm.io/gorm"
)

func subDistrictMigrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.SubDistricts{}) {
		db.Migrator().CreateTable(&models.SubDistricts{})
		fmt.Println("Success Create Table Sub Districts")
	}
}
