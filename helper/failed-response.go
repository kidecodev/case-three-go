package helper

import "github.com/gin-gonic/gin"

//ResponseFailed func
func ResponseFailed(str interface{}) gin.H {
	return gin.H{
		"error":   true,
		"message": str,
		"data":    nil,
	}
}

//ResponseFailedDataNil func
func ResponseFailedDataNil(str interface{}) gin.H {
	return gin.H{
		"error":   true,
		"message": str,
		"data":    nil,
	}
}
