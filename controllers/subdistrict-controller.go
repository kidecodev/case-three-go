package controllers

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetListSubDistrict func
func GetListSubDistrict(context *gin.Context) {
	db := *config.GetConnection()
	var (
		subdistrict []models.SubDistricts
		result      gin.H
	)

	db.Find(&subdistrict)
	if length := len(subdistrict); length <= 0 {
		result = helper.ResponseApi(subdistrict, length)
	} else {
		result = helper.ResponseApi(subdistrict, length)
	}

	context.JSON(http.StatusOK, result)
}

//GetSubDistrictByID func
func GetSubDistrictByID(context *gin.Context) {
	db := *config.GetConnection()
	var (
		subdistrict []models.SubDistricts
		result      gin.H
	)

	db.Find(&subdistrict)
	if length := len(subdistrict); length <= 0 {
		result = helper.ResponseApi(subdistrict, length)
	} else {
		result = helper.ResponseApi(subdistrict, length)
	}

	context.JSON(http.StatusOK, result)
}

//AddSubDistrict func
func AddSubDistrict(context *gin.Context) {
	db := *config.GetConnection()

	districtid, _ := strconv.ParseUint(context.PostForm("district_id"), 10, 8)

	subdistrict := models.SubDistricts{
		DistrictID: uint(districtid),
		Name:       context.PostForm("name"),
	}

	db.Create(&subdistrict)

	context.JSON(http.StatusOK, subdistrict.ID)
}

//UpdateSubDistrict func
func UpdateSubDistrict(context *gin.Context) {
	db := *config.GetConnection()
	subdistrict := &models.SubDistricts{}
	db.First(&subdistrict, context.Param("id"))

	districtid, _ := strconv.ParseUint(context.PostForm("district_id"), 10, 8)

	subdistrict.DistrictID = uint(districtid)
	subdistrict.Name = context.PostForm("name")

	db.Save(&subdistrict)

	context.JSON(http.StatusOK, subdistrict)
}

//DeleteSubDistrict func
func DeleteSubDistrict(context *gin.Context) {
	db := *config.GetConnection()

	db.Delete(&models.SubDistricts{}, context.Param("id"))

	context.JSON(http.StatusOK, nil)
}
