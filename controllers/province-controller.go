package controllers

import (
	"frmgol/helper"
	"frmgol/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListProvince func
func (strDB *StrDB) GetListProvince(c *gin.Context) {
	var (
		province []models.Provinces
		result   gin.H
	)

	strDB.DB.Find(&province)
	if length := len(province); length <= 0 {
		result = helper.ResponseApi(province, length)
	} else {
		result = helper.ResponseApi(province, length)
	}

	c.JSON(http.StatusOK, result)
}
