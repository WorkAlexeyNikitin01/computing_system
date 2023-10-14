package httpgin

import (
	"fmt"
	"lab/orderService/internal/order"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	Name     string  `json:"name"`
	Code     string	 `json:"code"`
	Quantity int	 `json:"quantity"`
	Price    float64 `json:"price"`
}

type OrderSuccess struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Code     string  `json:"code"`
	Price    float64 `json:"price"`
	Total    float64 `json:"total"`
}

func orderSuccess(or *order.Order) *gin.H {
	return &gin.H{
		"data": OrderSuccess{
			Id: fmt.Sprint(or.Id),
			Name: or.Name,
			Quantity: or.Quantity,
			Price: or.Price,
			Total: float64(or.Quantity)*or.Price,
			Code: or.Code,
		},
		"error": nil,
	}
}

func orderError(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}
