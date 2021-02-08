package controllers

import (
	"frmgol/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetReportCount ...
func GetReportCount(ctx *gin.Context)  {

}

//GetReportCity ...
func GetReportCity(ctx *gin.Context)  {
	result := service.GetReportCity(ctx)

	ctx.JSON(http.StatusOK, result)
}