package controllers

import (
	"frmgol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListSubDistrict func
func GetListSubDistrict(context *gin.Context) {
	result := service.GetListSubDistrict(context)

	context.JSON(http.StatusOK, result)
}

//GetSubDistrictByID func
func GetSubDistrictByID(context *gin.Context) {
	result := service.GetSubDistrictByID(context.Param("id"))

	context.JSON(http.StatusOK, result)
}

//AddSubDistrict func
func AddSubDistrict(context *gin.Context) {
	result := service.AddSubDistrict(context)

	context.JSON(http.StatusOK, result)
}

//UpdateSubDistrict func
func UpdateSubDistrict(context *gin.Context) {
	result := service.UpdateSubDistrict(context)

	context.JSON(http.StatusOK, result)
}

//DeleteSubDistrict func
func DeleteSubDistrict(context *gin.Context) {
	result := service.DeleteSubDistrict(context)

	context.JSON(http.StatusOK, result)
}
