package migration

import "frmgol/config"

//Migrations function
func Migrations() {
	db := *config.GetConnection()
	ProvinceMigrate(&db)
	DistrictMigrate(&db)
	SubDistrictMigrate(&db)
	PersonMigrate(&db)
	AddPhotoProfilToPerson(&db)
}
