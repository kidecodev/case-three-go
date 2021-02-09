package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"github.com/gin-gonic/gin"
)

//GetReportTotal
func GetReportTotal(context *gin.Context) interface{}  {
	db := *config.GetConnection()
	type response struct {
		Male   int64
		Female int64
	}

	var (
		result response
	)


	db.Raw("select * from" +
		"(" +
		"select count(*) as male from persons group by persons.gender HAVING persons.gender = 'M'" +
		") as m " +
		"NATURAL JOIN " +
		"(" +
		"select count(*) as female from persons group by persons.gender HAVING persons.gender = 'F'" +
		") as f").
		Scan(&result)

	return helper.ResponseSuccessSingle(&result)
}

//GetReportCity ...
func GetReportCity(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		persons   models.Persons
		location  string
		response []helper.ReportCityResponse
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
		var locations []string
		rows, _ := db.Table("persons").
			Select("d.name as location").
			Where("persons.id = ?", response[i].ID).
			Joins("left join person_handle_office_locations opl on opl.person_id = persons.id").
			Joins("left join office_locations o on opl.office_location_id = o.id").
			Joins("left join sub_districts sd on o.sub_district_id = sd.id").
			Joins("left join districts d on sd.district_id = d.id").
			Group("d.name").
			Rows()

		for rows.Next() {
			rows.Scan(&location)
			locations = append(locations, location)
		}
		response[i].Locations = locations
	}

	return helper.ResponseSuccessSingle(&response)
}

//GetReportByGender ...
func GetReportByGender(context *gin.Context) interface{} {
	db := *config.GetConnection()
	type response struct {
		Name   string
		Total  int64
		Male   int64
		Female int64
	}

	var (
		result []response
	)

	db.Raw("select * from (" +
		"select d.name, count(*) as total from persons " +
		"left join person_handle_office_locations opl on opl.person_id = persons.id " +
		"left join office_locations o on opl.office_location_id = o.id " +
		"left join sub_districts sd on o.sub_district_id = sd.id " +
		"left join districts d on sd.district_id = d.id " +
		"group by d.name" +
		") as c " +
		"NATURAL JOIN " +
		"(" +
		"select d.name, count(*) as male from persons " +
		"left join person_handle_office_locations opl on opl.person_id = persons.id " +
		"left join office_locations o on opl.office_location_id = o.id " +
		"left join sub_districts sd on o.sub_district_id = sd.id " +
		"left join districts d on sd.district_id = d.id " +
		"group by d.name, persons.gender HAVING persons.gender = 'M'" +
		") as m " +
		"NATURAL JOIN " +
		"(" +
		"select d.name, count(*) as female from persons " +
		"left join person_handle_office_locations opl on opl.person_id = persons.id " +
		"left join office_locations o on opl.office_location_id = o.id " +
		"left join sub_districts sd on o.sub_district_id = sd.id " +
		"left join districts d on sd.district_id = d.id " +
		"group by d.name, persons.gender HAVING persons.gender = 'F'" +
		") as f").
		Scan(&result)

	return helper.ResponseSuccessSingle(&result)
}
