package models

import "gorm.io/gorm"

//Districts definition
type Districts struct {
	gorm.Model
	ProvinceID uint
	Name       string
	SubDistrict []SubDistricts `gorm:"ForeignKey:DistrictID"`
}
