package migration

import (
	"frmgol/models"

	"gorm.io/gorm"
)

//Migrations function
func Migrations(db *gorm.DB) {
	db.Migrator().HasTable(&models.Provinces{})
}
