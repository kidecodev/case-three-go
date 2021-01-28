package web

import (
	"github.com/gin-gonic/gin"
)

//Route function
func Route(router *gin.Engine) *gin.Engine {

	router.GET("/total", controllers.GetListProvince)

	return router
}
