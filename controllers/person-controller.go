package controllers

import (
	"frmgol/config"
	"frmgol/helper"
	"frmgol/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetListPersons func
func GetListPersons(context *gin.Context) {
	db := *config.GetConnection()
	var (
		persons []models.Persons
		result  gin.H
	)

	db.Find(&persons)
	if length := len(persons); length <= 0 {
		result = helper.ResponseApi(persons, length)
	} else {
		result = helper.ResponseApi(persons, length)
	}

	context.JSON(http.StatusOK, result)
}

//GetPersonByID func
func GetPersonByID(context *gin.Context) {
	db := *config.GetConnection()
	var (
		person []models.Persons
		result gin.H
	)

	db.Find(&person)
	if length := len(person); length <= 0 {
		result = helper.ResponseApi(person, length)
	} else {
		result = helper.ResponseApi(person, length)
	}

	context.JSON(http.StatusOK, result)
}

//AddPerson func
func AddPerson(context *gin.Context) {
	db := *config.GetConnection()

	person := models.Persons{
		FullName: context.PostForm("name"),
	}

	db.Create(&person)

	context.JSON(http.StatusOK, person.ID)
}

//UpdatePerson func
func UpdatePerson(context *gin.Context) {
	db := *config.GetConnection()
	person := &models.Persons{}
	db.First(&person, context.Param("id"))

	person.FullName = context.PostForm("name")

	db.Save(&person)

	context.JSON(http.StatusOK, person)
}

//DeletePerson func
func DeletePerson(context *gin.Context) {
	db := *config.GetConnection()

	db.Delete(&models.Persons{}, context.Param("id"))

	context.JSON(http.StatusOK, nil)
}
