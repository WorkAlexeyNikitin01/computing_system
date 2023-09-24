package httpgin

import (
	"lab/storeroomService/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppStoreroomInterface) {
	r.POST("/add-product-to-store", AddToStoreroom(a))
	r.GET("/check-product-to-store", GetProductFromStoreroom(a))
	r.DELETE("/delete-product-to-store", DeleteFromStoreroom(a))
}