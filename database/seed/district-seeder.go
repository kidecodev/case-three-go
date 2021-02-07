package seed

import (
	"encoding/json"
	"frmgol/models"
	"frmgol/service"
	"frmgol/utils"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type rajaOngkirDistrictResponse struct {
	RajaOngkir districtResponse `json:"rajaongkir"`
}

type districtResponse struct {
	Results []resultDistrictResponse `json:"results"`
}

type resultDistrictResponse struct {
	CityID string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province string `json:"province"`
	Type string `json:"type"`
	CityName string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

func districtsSeeder(db *gorm.DB) {
	db.Exec("DELETE FROM districts")
	data := service.GetDataFromRajaOngkir(utils.GetEnv("RAJAONGKIR_API_URL") + "/city?key=" + utils.GetEnv("RAJAONGKIR_API_KEY"))
	var (
		response rajaOngkirDistrictResponse
		district models.Districts
	)

	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatalln("Error fetch response")
	}

	for i := 0; i < len(response.RajaOngkir.Results); i++ {
		provinceID, _ := strconv.Atoi(response.RajaOngkir.Results[i].ProvinceID)
		id, _ := strconv.Atoi(response.RajaOngkir.Results[i].CityID)
		district.ID = uint(id)
		district.ProvinceID = uint(provinceID)
		district.Name = response.RajaOngkir.Results[i].CityName
		db.Create(&district)
	}
}
