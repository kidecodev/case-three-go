package controllers

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListProvince func
func GetListProvince(c *gin.Context) {
	db := *config.GetConnection()
	var (
		province []models.Provinces
		result   gin.H
	)

	db.Find(&province)
	if length := len(province); length <= 0 {
		result = helper.ResponseApi(province, length)
	} else {
		result = helper.ResponseApi(province, length)
	}

	c.JSON(http.StatusOK, result)
}

//AddProvince func
func AddProvince(context *gin.Context) {
	db := *config.GetConnection()

	province := models.Provinces{Name: context.PostForm("name")}

	result := db.Create(&province)

	context.JSON(http.StatusOK, result)
}

//UpdateProvince func
func UpdateProvince(context *gin.Context) {
	db := *config.GetConnection()
	province := &models.Provinces{}
	db.First(&province, context.Param("id"))

	province.Name = context.PostForm("name")

	db.Save(&province)

	context.JSON(http.StatusOK, province)
}

//DeleteProvince func
func DeleteProvince(context *gin.Context) {
	db := *config.GetConnection()

	db.Delete(&models.Provinces{}, context.Param("id"))
	context.JSON(http.StatusOK, nil)
}
