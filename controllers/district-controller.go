package controllers

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetListDistrict func
func GetListDistrict(context *gin.Context) {
	db := *config.GetConnection()
	var (
		district []models.Districts
		result   gin.H
	)

	db.Find(&district)
	if length := len(district); length <= 0 {
		result = helper.ResponseApi(district, length)
	} else {
		result = helper.ResponseApi(district, length)
	}

	context.JSON(http.StatusOK, result)
}

//GetDistrictByID func
func GetDistrictByID(context *gin.Context) {
	db := *config.GetConnection()
	var (
		district []models.Districts
		result   gin.H
	)

	db.Find(&district)
	if length := len(district); length <= 0 {
		result = helper.ResponseApi(district, length)
	} else {
		result = helper.ResponseApi(district, length)
	}

	context.JSON(http.StatusOK, result)
}

//AddDistrict func
func AddDistrict(context *gin.Context) {
	db := *config.GetConnection()

	provinceid, _ := strconv.ParseUint(context.PostForm("province_id"), 10, 8)

	district := models.Districts{
		ProvinceID: uint(provinceid),
		Name:       context.PostForm("name"),
	}

	db.Create(&district)

	context.JSON(http.StatusOK, district.ID)
}

//UpdateDistrict func
func UpdateDistrict(context *gin.Context) {
	db := *config.GetConnection()
	district := &models.Districts{}
	db.First(&district, context.Param("id"))

	provinceid, _ := strconv.ParseUint(context.PostForm("province_id"), 10, 8)

	district.ProvinceID = uint(provinceid)
	district.Name = context.PostForm("name")

	db.Save(&district)

	context.JSON(http.StatusOK, district)
}

//DeleteDistrict func
func DeleteDistrict(context *gin.Context) {
	db := *config.GetConnection()

	db.Delete(&models.Districts{}, context.Param("id"))

	context.JSON(http.StatusOK, nil)
}
