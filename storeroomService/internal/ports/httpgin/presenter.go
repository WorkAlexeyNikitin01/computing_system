package httpgin

import "github.com/gin-gonic/gin"

type ProductCodeReq struct {
	code string
}

func storeroomError(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}

func storeroomSuccess(p interface{}) *gin.H {
	return &gin.H{
		"data": p,
		"error": nil,
	}
}
