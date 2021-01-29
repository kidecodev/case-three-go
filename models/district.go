package models

//Districts definition
type Districts struct {
	ID         uint `gorm:"primaryKey"`
	ProvinceID uint
	Name       string
}
