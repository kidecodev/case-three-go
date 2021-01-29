package web

import (
	"frmgol/controllers"

	"github.com/gin-gonic/gin"
)

//Route function
func Route(router *gin.Engine) *gin.Engine {
	router.GET("/province", controllers.GetListProvince)
	router.POST("/province", controllers.AddProvince)
	router.PUT("/province", controllers.UpdateProvince)
	router.DELETE("/province", controllers.DeleteProvince)

	router.GET("/district", controllers.GetListDistrict)
	router.POST("/district", controllers.AddDistrict)
	router.PUT("/district", controllers.UpdateDistrict)
	router.DELETE("/district", controllers.DeleteDistrict)

	return router
}
