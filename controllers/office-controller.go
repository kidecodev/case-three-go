package controllers

import (
	"frmgol/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetListOffice func
func GetListOffice(context *gin.Context) {
	result := service.GetListOffice(context)

	context.JSON(http.StatusOK, result)
}

//GetOfficeByID func
func GetOfficeByID(context *gin.Context) {
	result := service.GetOfficeByID(context.Param("id"))

	context.JSON(http.StatusOK, result)
}

//AddOffice func
func AddOffice(context *gin.Context) {
	result := service.AddOffice(context)

	context.JSON(http.StatusOK, result)
}

//UpdateOffice func
func UpdateOffice(context *gin.Context) {
	result := service.UpdateOffice(context)

	context.JSON(http.StatusOK, result)
}

//DeleteOffice func
func DeleteOffice(context *gin.Context) {
	result := service.DeleteOffice(context)

	context.JSON(http.StatusOK, result)
}