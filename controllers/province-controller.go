package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListProvince func
func GetListProvince(c *gin.Context) {
	// db := *config.GetConnection()
	// var (
	// 	province []models.Provinces
	// 	result   gin.H
	// )

	// &db.DB.Find(&province)
	// if length := len(province); length <= 0 {
	// 	result = helper.ResponseApi(province, length)
	// } else {
	// 	result = helper.ResponseApi(province, length)
	// }

	c.JSON(http.StatusOK, nil)
}

//AddProvince func
func AddProvince(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

//UpdateProvince func
func UpdateProvince(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}

//DeleteProvince func
func DeleteProvince(context *gin.Context) {
	context.JSON(http.StatusOK, nil)
}
