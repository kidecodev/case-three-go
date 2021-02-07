package models

type OfficePersonLocations struct {
	ID uint
	OfficeLocationID uint
	PersonID uint
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (OfficePersonLocations) TableName() string {
	return "person_handle_office_locations"
}