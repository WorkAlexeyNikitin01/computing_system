package service

import (
	"lab/first_microservice/internal/product"

	"github.com/gin-gonic/gin"
)


type productRequest struct {
	code        string
	name        string
	price       float64
	description string
}

type ProductResponse struct {
	id int
	code string
	price float64
	name string
	description  string
}

func productErr(err error) *gin.H {
	return &gin.H{
		"error": err,
	}
}

func productSuccess(p *product.Product) *gin.H {
	return &gin.H{
			"data": ProductResponse{
				code: p.Code,
				name: p.Name,
				description: p.Description,
				price: p.Price,
				id: p.Id,
			},
			"error": nil,
	}
}

func productsSuccess(p []*product.Product) {
	for _, i := range p {
		productSuccess(i)
	}
}
