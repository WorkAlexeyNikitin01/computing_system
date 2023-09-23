package httpgin

import (
	"lab/storeroomService/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppStoreroomInterface) {
	r.POST("/add-product", AddToStoreroom(a))
	r.GET("/check-product", GetProductFromStoreroom(a))
	r.DELETE("/delete-product", DeleteFromStoreroom(a))
}