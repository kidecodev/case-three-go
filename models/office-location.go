package models

import "gorm.io/gorm"

type OfficeLocations struct {
	gorm.Model
	SubDistrictID uint
	Name string
}