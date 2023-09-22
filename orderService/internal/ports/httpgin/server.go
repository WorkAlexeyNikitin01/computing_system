package httpgin

import (
	"lab/orderService/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(port string, a app.AppOrderInterface) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	s := &http.Server{
		Addr: port,
		Handler: router,
	}

	AppRouter(router.Group("api/v1"), a)

	return s
}
