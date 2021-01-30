package helper

//DistrictResponse struct
type DistrictResponse struct {
	ID       string
	Name     string
	Province string
}

//SubDistrictResponse struct
type SubDistrictResponse struct {
	ID       string
	Name     string
	District string
	Province string
}

//PersonResponse struct
type PersonResponse struct {
	ID           string
	Nip          string
	FullName     string
	FirstName    string
	LastName     string
	BirthDate    string
	BirthPlace   string
	Gender       string
	ZoneLocation string
	SubDistrict  string
	District     string
	Province     string
}
