package web

import (
	"gogin/controller"

	"github.com/gin-gonic/gin"
)

//Route function
func Route(router *gin.Engine) *gin.Engine {

	router.GET("/init-data", controller.InitItem)
	router.GET("/item", controller.GetItem)
	router.POST("/item", controller.AddItem)
	router.PUT("/item/:name", controller.UpdateItem)
	router.GET("/total", controller.TotalItem)

	return router
}
