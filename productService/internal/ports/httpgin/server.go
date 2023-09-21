package service

import (
	"lab/productService/internal/app"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(port string, a app.AppProductInterface) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	s := &http.Server{
		Addr: port,
		Handler: router,
	}

	AppRouter(router.Group("api/v1"), a)

	return s
}
