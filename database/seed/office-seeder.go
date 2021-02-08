package seed

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
)

func officeSeeder(db *gorm.DB)  {
	officeTemp := [...][3]interface{}{
		{1, 6, "Pengadilan Tanah Grogot"},
		{2, 6, "Kejaksaan Tanah Grogot"},
		{3, 6, "Polres Paser"},
	}

	var office models.OfficeLocations

	for _, value := range officeTemp {
		office.ID = uint(value[0].(int))
		office.Name = fmt.Sprintf("%v", value[2])
		office.SubDistrictID = uint(value[1].(int))
		db.Create(&office)
	}
}