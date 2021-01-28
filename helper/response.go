package helper

import "github.com/gin-gonic/gin"

//ResponseApi func
func ResponseApi(str interface{}, length int) gin.H {
	if length > 0 {
		return gin.H{
			"message": "success",
			"data":    str,
			"count":   length,
		}
	} else {
		return gin.H{
			"message": "success",
			"data":    nil,
			"count":   length,
		}
	}
}
