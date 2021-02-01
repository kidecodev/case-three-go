package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"frmgol/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetListPersons func
func GetListPersons(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		persons  []models.Persons
		response []helper.PersonResponse
	)

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "3"))

	paginator := utils.PagingCustomResult(&utils.Param{
		DB:      &db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
		Join:    "left join sub_districts on sub_districts.id = persons.sub_district_id left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id",
		Query:   "persons.*, sub_districts.name as sub_district, provinces.name as province, districts.name as district",
	}, &persons, &response)

	return helper.ResponseSuccessList(paginator)
}

//GetPersonByID func
func GetPersonByID(id interface{}) interface{} {
	db := *config.GetConnection()
	var (
		person   models.Persons
		response helper.PersonResponse
	)

	result := db.Model(&person).Select("persons.*, sub_districts.name as sub_district, provinces.name as province, districts.name as district").Joins("left join sub_districts on sub_districts.id = persons.sub_district_id left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id").Scan(&response)

	if result.RowsAffected == 0 {
		return helper.ResponseFailed(result.Error)
	}

	return helper.ResponseSuccessSingle(&response)
}

//AddPerson func
func AddPerson(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		subdistrict models.SubDistricts
		response    helper.PersonResponse
	)

	subdistrictid, _ := strconv.ParseUint(context.PostForm("sub_district_id"), 10, 8)

	result := db.First(&subdistrict, subdistrictid)

	if result.RowsAffected > 0 {
		person := models.Persons{
			SubDistrictID: uint(subdistrictid),
			Nip:           context.PostForm("nip"),
			FullName:      context.PostForm("full_name"),
			FirstName:     context.PostForm("first_name"),
			LastName:      context.PostForm("last_name"),
			BirthDate:     context.PostForm("birth_date"),
			BirthPlace:    context.PostForm("birth_place"),
			Gender:        context.PostForm("gender"),
			ZoneLocation:  context.PostForm("zone"),
		}

		resultCreate := db.Create(&person)

		if resultCreate.RowsAffected > 0 {
			db.Model(&person).Select("persons.*, sub_districts.name as sub_district, provinces.name as province, districts.name as district").Joins("left join sub_districts on sub_districts.id = persons.sub_district_id left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id").Scan(&response)

			return helper.ResponseSuccessSingle(response)
		}

		return helper.ResponseFailed(resultCreate.Error)
	}

	return helper.ResponseFailed(result.Error)
}

//UpdatePerson func
func UpdatePerson(context *gin.Context) interface{} {
	db := *config.GetConnection()
	var (
		subdistrict models.SubDistricts
		person      models.Persons
		response    helper.PersonResponse
	)

	subdistrictid, _ := strconv.ParseUint(context.PostForm("sub_district_id"), 10, 8)

	result := db.First(&subdistrict, subdistrictid)

	if result.RowsAffected > 0 {
		result := db.First(&person, context.Param("id"))

		if result.RowsAffected > 0 {

			person.SubDistrictID = uint(subdistrictid)
			person.Nip = context.PostForm("nip")
			person.FullName = context.PostForm("full_name")
			person.FirstName = context.PostForm("first_name")
			person.LastName = context.PostForm("last_name")
			person.BirthDate = context.PostForm("birth_date")
			person.BirthPlace = context.PostForm("birth_place")
			person.Gender = context.PostForm("gender")
			person.ZoneLocation = context.PostForm("zone")

			result := db.Save(&person)

			if result.RowsAffected > 0 {
				db.Model(&person).Select("persons.*, sub_districts.name as sub_district, provinces.name as province, districts.name as district").Joins("left join sub_districts on sub_districts.id = persons.sub_district_id left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id").Scan(&response)
				return helper.ResponseSuccessSingle(&response)
			}
		}
	}

	return helper.ResponseFailed(result.Error)
}

//UploadPhoto func
func UploadPhoto(context *gin.Context) interface{} {
	db := *config.GetConnection()
	var (
		person   models.Persons
		response helper.PersonResponse
	)

	result := db.First(&person, context.Param("id"))

	if result.RowsAffected > 0 {
		file, _ := context.FormFile("photo")
		path := "images/" + file.Filename

		if err := context.SaveUploadedFile(file, path); err != nil {
			return helper.ResponseFailed(err.Error())
		}

		person.ProfilUrl = path

		result := db.Save(&person)

		if result.RowsAffected > 0 {
			db.Model(&person).Select("persons.*, sub_districts.name as sub_district, provinces.name as province, districts.name as district").Joins("left join sub_districts on sub_districts.id = persons.sub_district_id left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id").Scan(&response)
			return helper.ResponseSuccessSingle(&response)
		}
	}

	return helper.ResponseFailed(result.Error)
}

//DeletePerson func
func DeletePerson(context *gin.Context) interface{} {
	db := *config.GetConnection()

	result := db.Delete(&models.Persons{}, context.PostForm("id"))

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessDelete("Success delete data person")
	}

	return helper.ResponseFailed(result.Error)
}
