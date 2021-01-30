package controllers

import (
	"frmgol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListDistrict func
func GetListDistrict(context *gin.Context) {
	result := service.GetListDistrict(context)

	context.JSON(http.StatusOK, result)
}

//GetDistrictByID func
func GetDistrictByID(context *gin.Context) {
	result := service.GetDistrictByID(context.Param("id"))

	context.JSON(http.StatusOK, result)
}

//AddDistrict func
func AddDistrict(context *gin.Context) {
	result := service.AddDistrict(context)

	context.JSON(http.StatusOK, result)
}

//UpdateDistrict func
func UpdateDistrict(context *gin.Context) {
	result := service.UpdateDistrict(context)

	context.JSON(http.StatusOK, result)
}

//DeleteDistrict func
func DeleteDistrict(context *gin.Context) {
	result := service.DeleteDistrict(context)

	context.JSON(http.StatusOK, result)
}
