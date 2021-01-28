package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListDistrict func
func GetListDistrict(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

//AddDistrict func
func AddDistrict(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

//UpdateDistrict func
func UpdateDistrict(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

//DeleteDistrict func
func DeleteDistrict(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}
