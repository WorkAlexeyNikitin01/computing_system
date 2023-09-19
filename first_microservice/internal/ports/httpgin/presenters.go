package service

import (
	"github.com/gin-gonic/gin"
)


type productRequest struct {
	code        string
	name        string
	price       float64
	description string
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
