package models

import "gorm.io/gorm"

// Provinces struct
type Provinces struct {
	gorm.Model
	Name string
	District []Districts `gorm:"ForeignKey:ProvinceID"`
}
