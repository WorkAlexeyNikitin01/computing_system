package httpgin

import (
	"lab/storeroomService/internal/storeroom"

	"github.com/gin-gonic/gin"
)

type ProductCodeReq struct {
	code     string
	quantity int
}

type ProductSuccess struct {
	code     string
	quantity int
}

func storeroomError(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}

func storeroomSuccess(p *storeroom.StoreroomProduct) *gin.H {
	return &gin.H{
		"data": ProductSuccess{
			code: p.CodeProduct,
			quantity: p.Quantity,
		},
		"error": nil,
	}
}
