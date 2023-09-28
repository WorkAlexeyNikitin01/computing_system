package httpgin

import (
	"lab/orderService/internal/app"
	store "lab/storeroomService/grpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(port string, a app.AppOrderInterface, client store.StoreroomServiceClient) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	s := &http.Server{
		Addr: port,
		Handler: router,
	}

	AppRouter(router.Group("api/v1"), a, client)

	return s
}
