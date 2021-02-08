package controllers

import (
	"frmgol/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//LoginUser ...
func LoginUser(ctx *gin.Context)  {
	result := service.AuthenticateUser(ctx)

	ctx.JSON(http.StatusOK, result)
}

//GetListUsers func
func GetListUsers(context *gin.Context) {
	result := service.GetListUsers(context)

	context.JSON(http.StatusOK, result)
}

//GetUserByID func
func GetUserByID(context *gin.Context) {
	result := service.GetUserByID(context.Param("id"))

	context.JSON(http.StatusOK, result)
}

//AddUser func
func AddUser(context *gin.Context) {
	result := service.AddUser(context)

	context.JSON(http.StatusOK, result)
}

//UpdateUser func
func UpdateUser(context *gin.Context) {
	result := service.UpdateUser(context)

	context.JSON(http.StatusOK, result)
}

//DeleteUser func
func DeleteUser(context *gin.Context) {
	result := service.DeleteUser(context)

	context.JSON(http.StatusOK, result)
}