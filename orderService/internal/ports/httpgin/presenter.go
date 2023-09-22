package httpgin

import (
	"lab/orderService/internal/order"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	listProduct []interface{}
	totalPrice  float64
}

type OrderSuccess struct {
	listProduct []interface{}
	totalPrice  float64
}

func orderSuccess(or *order.Order) *gin.H {
	return &gin.H{
		"data": OrderSuccess{
			listProduct: or.ListProducts,
			totalPrice: or.TotalPrice,
		},
		"error": nil,
	}
}

func orderError(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}
