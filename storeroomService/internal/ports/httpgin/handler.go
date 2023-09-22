package httpgin

import (
	"lab/storeroomService/internal/app"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductFromStoreroom(a app.AppStoreroomInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductCodeReq
		err := c.Bind(&req)
		if err != nil {
			c.JSON(200, storeroomError(err))
			return
		}

		result, err := a.GetProductFromStoreroom(req.code)
		if err != nil {
			c.JSON(200, storeroomError(err))
			return
		}
		log.Println(http.StatusOK, result)
		storeroomSuccess(result)
		c.JSON(200, result)
	}
}