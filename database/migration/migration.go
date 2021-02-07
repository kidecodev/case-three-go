package migration

import (
	"frmgol/config"
	"frmgol/utils"
)

//Migrations function
func Migrations() {
	db := *config.GetConnection()

	if utils.GetEnv("USE_MIGRATE") == "yes" {
		provinceMigrate(&db)
		districtMigrate(&db)
		subDistrictMigrate(&db)
		personMigrate(&db)
		addPhotoProfilToPerson(&db)
		personOfficeLocationMigrate(&db)
		officeLocationMigrate(&db)
		usersMigrate(&db)
	}

}
