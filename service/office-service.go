package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"frmgol/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//GetListOffice func
func GetListOffice(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		office []models.OfficeLocations
	)

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))

	paginator := utils.Paging(&utils.Param{
		DB:      &db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &office)

	return helper.ResponseSuccessList(paginator)
}

//GetOfficeByID func
func GetOfficeByID(id interface{}) interface{} {
	db := *config.GetConnection()

	var (
		office models.OfficeLocations
	)

	result := db.First(&office, id)

	if result.RowsAffected == 0 {
		return helper.ResponseFailed(result.Error)
	}

	return helper.ResponseSuccessSingle(&office)
}

//AddOffice func
func AddOffice(context *gin.Context) interface{} {
	db := *config.GetConnection()

	office := models.OfficeLocations{
		SubDistrictID: uint(1), //context.PostForm("name"),
		Name: context.PostForm("password"),
	}

	result := db.Create(&office)

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessSingle(&office)
	}

	return helper.ResponseFailed(result.Error)
}

//UpdateOffice func
func UpdateOffice(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		office models.OfficeLocations
	)

	result := db.First(&office, context.Param("id"))

	if result.RowsAffected > 0 {
		office.SubDistrictID = uint(1) //context.PostForm("sub_district_id")
		office.Name = context.PostForm("name")

		result := db.Save(&office)
		if result.RowsAffected > 0 {
			return helper.ResponseSuccessSingle(&office)
		}
	}

	return helper.ResponseFailed(result.Error)
}

//DeleteOffice func
func DeleteOffice(context *gin.Context) interface{} {
	db := *config.GetConnection()

	result := db.Delete(&models.OfficeLocations{}, context.PostForm("id"))

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessDelete("Success delete data office")
	}

	return helper.ResponseFailed(result.Error)
}