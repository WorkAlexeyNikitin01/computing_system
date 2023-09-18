package service

import (
	"lab/first_microservice/internal/app"
	
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allListProducts(a app.AppProductInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody productRequest
		err := c.Bind(&reqBody)
		if err != nil {
			log.Println("error create ad", err)
			c.JSON(400, productErr(err))
			return
		}

		products, err := a.AllListProducts()
		if err != nil {
			c.JSON(200, productErr(err))
			log.Println("error get list products")
			return
		}
		log.Println("Success list products", http.StatusOK, "id ad", len(products))
		c.Status(http.StatusOK)
		c.JSON(200, productSuccess(products))
		log.Default()
	}
}