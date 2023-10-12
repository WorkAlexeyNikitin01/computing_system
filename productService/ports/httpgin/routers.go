package httpgin

import (
	"lab/productService/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppProductInterface) {
	r.POST("/get-product", getProduct(a))
	r.GET("/products", allListProducts(a))
	r.POST("/create-product", createProduct(a))
	r.DELETE("/delete-product", deleteProduct(a))
}
