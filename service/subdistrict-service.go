package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"frmgol/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetListSubDistrict func
func GetListSubDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		subdistrict []models.SubDistricts
		response    []helper.SubDistrictResponse
	)

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "3"))

	paginator := utils.PagingCustomResult(&utils.Param{
		DB:      &db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
		Join:    "left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id",
		Query:   "sub_districts.id, sub_districts.name, districts.name as district, provinces.name as province",
	}, &subdistrict, &response)

	return helper.ResponseSuccessList(paginator)
}

//GetSubDistrictByID func
func GetSubDistrictByID(id interface{}) interface{} {
	db := *config.GetConnection()
	var (
		subdistrict models.SubDistricts
		response    helper.SubDistrictResponse
	)

	result := db.Model(&subdistrict).Select("sub_districts.id, sub_districts.name, districts.name as district, provinces.name as province").Joins("left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id").Scan(&response)

	if result.RowsAffected == 0 {
		return helper.ResponseFailed(result.Error)
	}

	return helper.ResponseSuccessSingle(&response)
}

//AddSubDistrict func
func AddSubDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		district models.Districts
		response helper.SubDistrictResponse
	)

	districtid, _ := strconv.ParseUint(context.PostForm("district_id"), 10, 8)

	result := db.First(&district, districtid)

	if result.RowsAffected > 0 {
		subdistrict := models.SubDistricts{
			DistrictID: uint(districtid),
			Name:       context.PostForm("name"),
		}

		resultCreate := db.Create(&subdistrict)

		if resultCreate.RowsAffected > 0 {
			db.Model(&subdistrict).Select("sub_districts.id, sub_districts.name, provinces.name as province, districts.name as district").Joins("left join districts on districts.id = sub_districts.district_id").Joins("left join provinces on provinces.id = districts.province_id").Scan(&response)

			return helper.ResponseSuccessSingle(response)
		}

		return helper.ResponseFailed(resultCreate.Error)
	}

	return helper.ResponseFailed(result.Error)
}

//UpdateSubDistrict func
func UpdateSubDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()
	var (
		subdistrict models.SubDistricts
		district    models.Districts
		response    helper.SubDistrictResponse
	)

	districtid, _ := strconv.ParseUint(context.PostForm("district_id"), 10, 8)

	result := db.First(&district, districtid)

	if result.RowsAffected > 0 {
		result := db.First(&subdistrict, context.Param("id"))

		if result.RowsAffected > 0 {

			subdistrict.DistrictID = uint(districtid)
			subdistrict.Name = context.PostForm("name")

			result := db.Save(&subdistrict)
			if result.RowsAffected > 0 {
				db.Model(&subdistrict).Select("sub_districts.id, sub_districts.name, provinces.name as province, districts.name as district").Joins("left join districts on districts.id = sub_districts.district_id left join provinces on provinces.id = districts.province_id").Scan(&response)
				return helper.ResponseSuccessSingle(&response)
			}
		}
	}

	return helper.ResponseFailed(result.Error)
}

//DeleteSubDistrict func
func DeleteSubDistrict(context *gin.Context) interface{} {
	db := *config.GetConnection()

	result := db.Delete(&models.SubDistricts{}, context.PostForm("id"))

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessDelete("Success delete data sub district")
	}

	return helper.ResponseFailed(result.Error)
}
