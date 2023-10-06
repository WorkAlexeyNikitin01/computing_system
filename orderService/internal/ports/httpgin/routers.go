package httpgin

import (
	"lab/orderService/internal/app"
	store "lab/storeroomService/grpc"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func AppRouter(r *gin.RouterGroup, a app.AppOrderInterface, client store.StoreroomServiceClient, rabbitChan *amqp091.Channel) {
	r.POST("/create-order", createOrder(a, client, rabbitChan))
}
