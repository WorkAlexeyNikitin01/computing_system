package httpgin

import (
	"lab/orderService/internal/app"
	store "lab/storeroomService/grpc"
	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.AppOrderInterface, client store.StoreroomServiceClient) {
	r.POST("/create-order", createOrder(a, client))
}
