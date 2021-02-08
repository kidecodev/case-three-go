package seed

import (
	"frmgol/models"
	"gorm.io/gorm"
	"math/rand"
)

func personLocationSeeder(db *gorm.DB)  {
	var office models.OfficePersonLocations

	for i := 1; i <= 20; i++ {
		office.ID = uint(i)
		office.OfficeLocationID = uint(rand.Intn(2) + 1)
		office.PersonID = uint(rand.Intn(19) + 1)
		db.Create(&office)
	}
}