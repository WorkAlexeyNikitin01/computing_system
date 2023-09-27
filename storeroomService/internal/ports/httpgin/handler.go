package httpgin

import (
	"lab/storeroomService/internal/app"
	"lab/storeroomService/internal/storeroom"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductFromStoreroom(a app.AppStoreroomInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductCodeReq
		err := c.Bind(&req)
		if err != nil {
			log.Fatal(err)
			return
		}
		result, err := a.GetFromStoreroom(req.Code)
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
		result, err := a.AddToStoreroom(storeroom.StoreroomProduct{Code: req.Code, Quantity: req.Quantity})
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
		result, err := a.DeleteFromStoreroom(req.Code)
		if err != nil {
			c.JSON(200, storeroomError(err))
			log.Fatal(err)
			return
		}
		c.JSON(200, storeroomSuccess(result))
		log.Println("success delete from database", result)
	}
}
