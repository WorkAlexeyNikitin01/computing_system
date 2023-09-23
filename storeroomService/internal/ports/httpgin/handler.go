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

		result, err := a.GetFromStoreroom(req.code)
		if err != nil {
			c.JSON(200, storeroomError(err))
			return
		}
		log.Println(http.StatusOK, result)
		storeroomSuccess(result)
		c.JSON(200, result)
	}
}

func AddToStoreroom(a app.AppStoreroomInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductCodeReq
		err := c.Bind(&req)
		if err != nil {
			log.Fatal(err)
			return
		}
		result, err := a.AddToStoreroom(req.code, req.quantity)
		if err != nil {
			c.JSON(200, storeroomError(err))
			log.Fatal(err)
			return
		}
		c.JSON(200, storeroomSuccess(result))
	}
}

func DeleteFromStoreroom(a app.AppStoreroomInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductCodeReq
		err := c.Bind(&req)
		if err != nil {
			log.Fatal(err)
		}
		result, err := a.DeleteFromStoreroom(req.code)
		if err != nil {
			c.JSON(200, storeroomError(err))
			log.Fatal(err)
			return
		}
		c.JSON(200, storeroomSuccess(result))
	}
}
