package seed

import (
	"frmgol/config"
	"frmgol/utils"
)

//Seeder function
func Seeder() {
	db := *config.GetConnection()

	if utils.GetEnv("USE_SEEDER") == "yes" {
		provincesSeeder(&db)
		districtsSeeder(&db)
		subdistrictSeeder(&db)
		personSeeder(&db)
		officeSeeder(&db)
		personLocationSeeder(&db)
	}

}