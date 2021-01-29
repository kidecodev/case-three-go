package models

import (
	"gorm.io/gorm"
)

//SubDistricts struct
type SubDistricts struct {
	gorm.Model
	DistrictID uint
	Name       string
}
