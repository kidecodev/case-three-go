package web

import (
	"frmgol/controllers"

	"github.com/gin-gonic/gin"
)

//Route function
func Route(router *gin.Engine) *gin.Engine {
	router.GET("/province", controllers.GetListProvince)
	router.GET("/province/:id", controllers.GetProvinceByID)
	router.POST("/province", controllers.AddProvince)
	router.PUT("/province/:id", controllers.UpdateProvince)
	router.DELETE("/province", controllers.DeleteProvince)

	router.GET("/district", controllers.GetListDistrict)
	router.GET("/district/:id", controllers.GetDistrictByID)
	router.POST("/district", controllers.AddDistrict)
	router.PUT("/district/:id", controllers.UpdateDistrict)
	router.DELETE("/district", controllers.DeleteDistrict)

	router.GET("/subdistrict", controllers.GetListSubDistrict)
	router.GET("/subdistrict/:id", controllers.GetSubDistrictByID)
	router.POST("/subdistrict", controllers.AddSubDistrict)
	router.PUT("/subdistrict/:id", controllers.UpdateSubDistrict)
	router.DELETE("/subdistrict", controllers.DeleteSubDistrict)

	router.GET("/persons", controllers.GetListPersons)
	router.GET("/person/:id", controllers.GetPersonByID)
	router.POST("/person", controllers.AddPerson)
	router.PUT("/person/:id", controllers.UpdatePerson)
	router.DELETE("/person", controllers.DeletePerson)

	router.PUT("/person/:id/photo", controllers.UploadPhoto)

	return router
}
