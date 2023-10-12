package httpgin

import (
	"lab/productService/product"

	"github.com/gin-gonic/gin"
)


type productRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductResponse struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Price float64
	Name string
	Description  string
}

func productErr(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}

func productSuccess(p *product.Product) *gin.H {
	return &gin.H{
			"data": ProductResponse{
				Code: p.Code,
				Name: p.Name,
				Description: p.Description,
				Price: p.Price,
				Id: p.Id,
			},
			"error": nil,
	}
}

func productsSuccess(p []*product.Product) {
	for _, i := range p {
		productSuccess(i)
	}
}
