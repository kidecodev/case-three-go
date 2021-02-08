package seed

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
)

func subdistrictSeeder(db *gorm.DB) {
	db.Exec("DELETE FROM sub_districts CASCADE")

	subdistrictTemp := [...][3]interface{}{
		{1, 341, "BATU SOPANG"},
		{2, 341, "MUARA SAMU"},
		{3, 341, "BATU ENGAU"},
		{4, 341, "TANJUNG HARAPAN"},
		{5, 341, "PASIR BELENGKONG"},
		{6, 341, "TANAH GROGOT"},
		{7, 341, "KUARO"},
		{8, 341, "LONG IKIS"},
		{9, 341, "MUARA KOMAM"},
		{10, 341, "LONG KALI"},
		{11, 354, "PENAJAM"},
		{12, 354, "SEPAKU"},
	}

	var subDistrict models.SubDistricts

	for _, value := range subdistrictTemp {
		subDistrict.ID = uint(value[0].(int))
		subDistrict.Name = fmt.Sprintf("%v", value[2])
		subDistrict.DistrictID = uint(value[1].(int))
		db.Create(&subDistrict)
	}
}