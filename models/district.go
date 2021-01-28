package models

// District definition
type Districts struct {
	ID         uint `gorm:"primaryKey"`
	ProvinceID uint
	Name       string
}
