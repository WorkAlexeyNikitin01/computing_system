package httpgin

import (
	"lab/orderService/internal/app"

	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppOrderInterface) {
	r.POST("/order", createOrder(a))
}
