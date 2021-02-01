package controllers

import (
	"frmgol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListPersons func
func GetListPersons(context *gin.Context) {
	result := service.GetListPersons(context)

	context.JSON(http.StatusOK, result)
}

//GetPersonByID func
func GetPersonByID(context *gin.Context) {
	result := service.GetPersonByID(context.Param("id"))

	context.JSON(http.StatusOK, result)
}

//AddPerson func
func AddPerson(context *gin.Context) {
	result := service.AddPerson(context)

	context.JSON(http.StatusOK, result)
}

//UpdatePerson func
func UpdatePerson(context *gin.Context) {
	result := service.UpdatePerson(context)

	context.JSON(http.StatusOK, result)
}

//UploadPhoto func
func UploadPhoto(context *gin.Context) {
	result := service.UploadPhoto(context)

	context.JSON(http.StatusOK, result)
}

//DeletePerson func
func DeletePerson(context *gin.Context) {
	result := service.DeletePerson(context)

	context.JSON(http.StatusOK, result)
}
