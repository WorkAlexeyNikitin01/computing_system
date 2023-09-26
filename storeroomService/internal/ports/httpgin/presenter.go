package httpgin

import (
	"lab/storeroomService/internal/storeroom"

	"github.com/gin-gonic/gin"
)

type ProductCodeReq struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type ProductSuccess struct {
	Id int `json:"id"`
	Code     string `json:"code"`
	Quantity int `json:"quantity"`
}

func storeroomError(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}

func storeroomSuccess(p *storeroom.StoreroomProduct) *gin.H {
	return &gin.H{
		"data": ProductSuccess{
			Id: p.Id,
			Code: p.Code,
			Quantity: p.Quantity,
		},
		"error": nil,
	}
}
