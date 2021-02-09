package web

import (
	"frmgol/controllers"

	"github.com/gin-gonic/gin"
)

//Route function
func Route(router *gin.Engine) *gin.Engine {
	router.POST("/auth", controllers.LoginUser)

	router.GET("/province", controllers.GetListProvince)
	router.GET("/province/:id", controllers.GetProvinceByID)
	router.GET("/district", controllers.GetListDistrict)
	router.GET("/district/:id", controllers.GetDistrictByID)
	router.GET("/subdistrict", controllers.GetListSubDistrict)
	router.GET("/subdistrict/:id", controllers.GetSubDistrictByID)

	group := router.Group("/")

	group.Use(Middleware)
	{
		group.POST("/province", controllers.AddProvince)
		group.PUT("/province/:id", controllers.UpdateProvince)
		group.DELETE("/province", controllers.DeleteProvince)

		group.POST("/district", controllers.AddDistrict)
		group.PUT("/district/:id", controllers.UpdateDistrict)
		group.DELETE("/district", controllers.DeleteDistrict)

		group.POST("/subdistrict", controllers.AddSubDistrict)
		group.PUT("/subdistrict/:id", controllers.UpdateSubDistrict)
		group.DELETE("/subdistrict", controllers.DeleteSubDistrict)

		group.GET("/persons", controllers.GetListPersons)
		group.GET("/person/:id", controllers.GetPersonByID)
		group.POST("/person", controllers.AddPerson)
		group.PUT("/person/:id", controllers.UpdatePerson)
		group.DELETE("/person", controllers.DeletePerson)
		group.PUT("/person/:id/photo", controllers.UploadPhoto)

		// Add Route Office Location
		group.GET("/offices", controllers.GetListOffice)
		group.GET("/office/:id", controllers.GetOfficeByID)
		group.POST("/office", controllers.AddOffice)
		group.PUT("/office/:id", controllers.UpdateOffice)
		group.DELETE("/office", controllers.DeleteOffice)

		// Add Route Users
		group.GET("/users", controllers.GetListUsers)
		group.GET("/user/:id", controllers.GetUserByID)
		group.POST("/user", controllers.AddUser)
		group.PUT("/user/:id", controllers.UpdateUser)
		group.DELETE("/user", controllers.DeleteUser)

		// Add Route Report
		group.GET("/report/person/count", controllers.GetReportCount)
		group.GET("/report/person/city", controllers.GetReportCity)
		group.GET("/report/person/gender", controllers.GetReportByGender)
	}


	return router
}
