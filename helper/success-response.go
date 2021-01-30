package helper

import "github.com/gin-gonic/gin"

//ResponseSuccess func
func ResponseSuccess(str interface{}, length int) gin.H {
	return gin.H{
		"error":   false,
		"message": "success",
		"data":    str,
		"count":   length,
	}
}

//ResponseSuccessList func
func ResponseSuccessList(str interface{}) gin.H {
	return gin.H{
		"error":   false,
		"message": "success",
		"data":    str,
	}
}

//ResponseSuccessSingle func
func ResponseSuccessSingle(str interface{}) gin.H {
	return gin.H{
		"error":   false,
		"message": "success",
		"data":    str,
	}
}

//ResponseSuccessDelete func
func ResponseSuccessDelete(str interface{}) gin.H {
	return gin.H{
		"error":   false,
		"message": str,
	}
}
