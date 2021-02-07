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

type rajaOngkirProvinceResponse struct {
	RajaOngkir provinceResponse `json:"rajaongkir"`
}

type provinceResponse struct {
	Results []resultProvinceResponse `json:"results"`
}

type resultProvinceResponse struct {
	ProvinceID string `json:"province_id"`
	Province string `json:"province"`
}

func provincesSeeder(db *gorm.DB) {
	db.Exec("DELETE FROM provinces")
	data := service.GetDataFromRajaOngkir(utils.GetEnv("RAJAONGKIR_API_URL") + "/province?key=" + utils.GetEnv("RAJAONGKIR_API_KEY"))
	var (
		response rajaOngkirProvinceResponse
		province models.Provinces
	)

	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatalln("Error fetch response")
	}

	for i := 0; i < len(response.RajaOngkir.Results); i++ {
		id, _ := strconv.Atoi(response.RajaOngkir.Results[i].ProvinceID)
		province.ID = uint(id)
		province.Name = response.RajaOngkir.Results[i].Province
		db.Create(&province)
	}
}
