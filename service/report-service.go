package service

import (
	"fmt"
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"github.com/gin-gonic/gin"
)

//GetPersonTotal ...
func GetPersonTotal()  {

}

type resTemp struct {
	Name string
}

//GetReportCity ...
func GetReportCity(context *gin.Context) interface{}  {
	db := *config.GetConnection()

	var (
		persons models.Persons
		locations []string
		response    []helper.ReportCityResponse
	)

	db.Model(&persons).
		Select("persons.id, persons.full_name, count(*) as total").
		Joins("left join person_handle_office_locations opl on opl.person_id = persons.id").
		Joins("left join office_locations o on opl.office_location_id = o.id").
		Joins("left join sub_districts sd on o.sub_district_id = sd.id").
		Joins("left join districts d on sd.district_id = d.id").
		Group("persons.id").
		Scan(&response)

	for i := 0; i < len(response); i++ {
		db.Model(&persons).
			Select("d.name").
			Joins("left join person_handle_office_locations opl on opl.person_id = persons.id").
			Joins("left join office_locations o on opl.office_location_id = o.id").
			Joins("left join sub_districts sd on o.sub_district_id = sd.id").
			Joins("left join districts d on sd.district_id = d.id").
			Where("persons.id = ?", response[i].ID).
			Scan(&locations)
		fmt.Println(locations)
		//response[i].Locations = locations
	}


	return helper.ResponseSuccessSingle(&response)
}