package models

import "gorm.io/gorm"

//Persons structs
type Persons struct {
	gorm.Model
	SubDistrictID uint
	Nip           string
	FullName      string
	FirstName     string
	LastName      string
	BirthDate     string
	BirthPlace    string
	ProfilUrl     string
	Gender        string
	ZoneLocation  string

}
