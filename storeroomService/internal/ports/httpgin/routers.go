package httpgin

import (
	"lab/storeroomService/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppStoreroomInterface) {
	r.POST("/add-sproduct", AddToStoreroom(a))
	r.POST("/check-sproduct", GetProductFromStoreroom(a)) // ПОЧЕМУ НЕ GET? НЕ ОТПРАВЛЯЕТ JSON В hANDLER
	r.DELETE("/delete-sproduct", DeleteFromStoreroom(a))
}