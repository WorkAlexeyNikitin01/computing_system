package service

import (
	"lab/first_microservice/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppProductInterface) {
	r.GET("/products", allListProducts(a))
	r.POST("/create-product", createProduct(a))
}
