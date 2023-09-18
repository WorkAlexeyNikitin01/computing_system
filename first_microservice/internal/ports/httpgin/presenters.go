package service

import (
	"github.com/gin-gonic/gin"
)


type productRequest struct {
	nameProduct string
}

type ProductResponse struct {

}

func productErr(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}

func productSuccess(p interface{}) *gin.H {
	return &gin.H{
		"objects": p,
	}
}
