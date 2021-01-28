package migration

import (
	"gorm.io/gorm"
)

//Migrations function
func Migrations(db *gorm.DB) {
	ProvinceMigrate(db)
	DistrictMigrate(db)
}
