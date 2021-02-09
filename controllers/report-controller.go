package controllers

import (
	"frmgol/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetReportCount ...
func GetReportCount(ctx *gin.Context)  {
	result := service.GetReportTotal(ctx)

	ctx.JSON(http.StatusOK, result)
}

//GetReportCity ...
func GetReportCity(ctx *gin.Context)  {
	result := service.GetReportCity(ctx)

	ctx.JSON(http.StatusOK, result)
}

//GetReportByGender ...
func GetReportByGender(ctx *gin.Context)  {
	result := service.GetReportByGender(ctx)

	ctx.JSON(http.StatusOK, result)
}