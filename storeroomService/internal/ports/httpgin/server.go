package httpgin

import (
	"lab/storeroomService/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(port string, app app.AppStoreroomInterface) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	server := &http.Server{
		Addr: port,
		Handler: router,
	}
	AppRouter(router.Group("api/v1"), app)

	return server
}
