package migration

import (
	"fmt"
	"frmgol/models"

	"gorm.io/gorm"
)

func addPhotoProfilToPerson(db *gorm.DB) {
	if !db.Migrator().HasColumn(&models.Persons{}, "ProfilUrl") {
		db.Migrator().AddColumn(&models.Persons{}, "ProfilUrl")
		fmt.Println("Success Add Profil URL to table persons")
	}
}
