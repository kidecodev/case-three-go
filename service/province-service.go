package service

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"frmgol/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetListProvince func
func GetListProvince(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		province []models.Provinces
	)

	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))

	paginator := utils.Paging(&utils.Param{
		DB:      &db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &province)

	return helper.ResponseSuccessList(paginator)
}

//GetProvinceByID func
func GetProvinceByID(id interface{}) interface{} {
	db := *config.GetConnection()

	var (
		province models.Provinces
	)

	result := db.First(&province, id)

	if result.RowsAffected == 0 {
		return helper.ResponseFailed(result.Error)
	}

	return helper.ResponseSuccessSingle(&province)
}

//AddProvince func
func AddProvince(context *gin.Context) interface{} {
	db := *config.GetConnection()

	province := models.Provinces{Name: context.PostForm("name")}

	result := db.Create(&province)

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessSingle(&province)
	}

	return helper.ResponseFailed(result.Error)
}

//UpdateProvince func
func UpdateProvince(context *gin.Context) interface{} {
	db := *config.GetConnection()

	var (
		province models.Provinces
	)

	result := db.First(&province, context.Param("id"))

	if result.RowsAffected > 0 {
		province.Name = context.PostForm("name")

		result := db.Save(&province)
		if result.RowsAffected > 0 {
			return helper.ResponseSuccessSingle(&province)
		}
	}

	return helper.ResponseFailed(result.Error)
}

//DeleteProvince func
func DeleteProvince(context *gin.Context) interface{} {
	db := *config.GetConnection()

	result := db.Delete(&models.Provinces{}, context.PostForm("id"))

	if result.RowsAffected > 0 {
		return helper.ResponseSuccessDelete("Success delete data province")
	}

	return helper.ResponseFailed(result.Error)
}
