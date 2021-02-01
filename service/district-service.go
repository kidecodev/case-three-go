package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"frmgol/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetListDistrict func
func GetListDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		district []models.Districts
		response []helper.DistrictResponse
	)

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "3"))

	paginator := utils.PagingCustomResult(&utils.Param{
		DB:      &db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
		Join:    "left join provinces on provinces.id = districts.province_id",
		Query:   "districts.id, districts.name, provinces.name as province",
	}, &district, &response)

	return helper.ResponseSuccessList(paginator)
}

//GetDistrictByID func
func GetDistrictByID(id interface{}) interface{} {
	db := *config.GetConnection()
	var (
		district models.Districts
		response helper.DistrictResponse
	)

	result := db.Model(&district).Select("districts.id, districts.name, provinces.name as province").Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)

	if result.RowsAffected == 0 {
		return helper.ResponseFailed(result.Error)
	}

	return helper.ResponseSuccessSingle(&response)
}

//AddDistrict func
func AddDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		province models.Provinces
		response helper.DistrictResponse
	)

	provinceid, _ := strconv.ParseUint(context.PostForm("province_id"), 10, 8)

	result := db.First(&province, provinceid)

	if result.RowsAffected > 0 {
		district := models.Districts{
			ProvinceID: uint(provinceid),
			Name:       context.PostForm("name"),
		}

		resultCreate := db.Create(&district)

		if resultCreate.RowsAffected > 0 {
			db.Model(&district).Select("districts.id, districts.name, provinces.name as province").Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)
			return helper.ResponseSuccessSingle(response)
		}

		return helper.ResponseFailed(resultCreate.Error)
	}

	return helper.ResponseFailed(result.Error)
}

//UpdateDistrict func
func UpdateDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()
	var (
		province models.Provinces
		district models.Districts
		response helper.DistrictResponse
	)

	provinceid, _ := strconv.ParseUint(context.PostForm("province_id"), 10, 8)

	result := db.First(&province, provinceid)

	if result.RowsAffected > 0 {
		result := db.First(&district, context.Param("id"))

		if result.RowsAffected > 0 {

			district.ProvinceID = uint(provinceid)
			district.Name = context.PostForm("name")

			result := db.Save(&district)
			if result.RowsAffected > 0 {
				db.Model(&district).Select("districts.id, districts.name, provinces.name as province").
					Joins("left join provinces on provinces.id = districts.province_id").
					Scan(&response)
				return helper.ResponseSuccessSingle(&response)
			}
		}
	}

	return helper.ResponseFailed(result.Error)
}

//DeleteDistrict func
func DeleteDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()

	result := db.Delete(&models.Districts{}, context.PostForm("id"))

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessDelete("Success delete data district")
	}

	return helper.ResponseFailed(result.Error)
}
