package httpgin

import (
	"lab/storeroomService/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppStoreroomInterface) {
	r.GET("/check-product", GetProductFromStoreroom(a))
}