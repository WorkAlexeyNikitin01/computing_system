package httpgin

import (
	"lab/orderService/internal/app"
	store "lab/storeroomService/grpc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func NewHTTPServer(port string, a app.AppOrderInterface, client store.StoreroomServiceClient, rabbitChan *amqp091.Channel) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	s := &http.Server{
		Addr: port,
		Handler: router,
	}

	AppRouter(router.Group("api/v1"), a, client, rabbitChan)

	return s
}
