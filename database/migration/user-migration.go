package migration

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
)

func usersMigrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Users{}) {
		db.Migrator().CreateTable(&models.Users{})
		fmt.Println("Success Create Table Users")
	}
}