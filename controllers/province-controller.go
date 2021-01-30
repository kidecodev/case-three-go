package controllers

import (
	"frmgol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListProvince func
func GetListProvince(context *gin.Context) {
	result := service.GetListProvince(context)

	context.JSON(http.StatusOK, result)
}

//GetProvinceByID func
func GetProvinceByID(context *gin.Context) {
	result := service.GetProvinceByID(context.Param("id"))

	context.JSON(http.StatusOK, result)
}

//AddProvince func
func AddProvince(context *gin.Context) {
	result := service.AddProvince(context)

	context.JSON(http.StatusOK, result)
}

//UpdateProvince func
func UpdateProvince(context *gin.Context) {
	result := service.UpdateProvince(context)

	context.JSON(http.StatusOK, result)
}

//DeleteProvince func
func DeleteProvince(context *gin.Context) {
	result := service.DeleteProvince(context)

	context.JSON(http.StatusOK, result)
}
