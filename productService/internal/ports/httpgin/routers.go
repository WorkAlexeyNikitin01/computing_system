package httpgin

import (
	"lab/productService/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppProductInterface) {
	r.GET("/products", allListProducts(a))
	r.POST("/create-product", createProduct(a))
	r.DELETE("/delete-product", deleteProduct(a))
}
