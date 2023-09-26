package httpgin

import (
	"lab/orderService/internal/app"
	"lab/orderService/internal/order"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createOrder(a app.AppOrderInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody OrderRequest
		err := c.Bind(&reqBody)
		if err != nil {
			log.Println("error create order", err)
			c.JSON(400, orderError(err))
			return
		}

		order, err := a.CreateOrder(order.Order{
			Quantity: reqBody.Quantity,
			Name: reqBody.Name,
			Code: reqBody.Code,
			Price: reqBody.Price,
		})
		if err != nil {
			c.JSON(200, orderError(err))
			log.Println("error create order")
			return
		}
		log.Println("Success create order", http.StatusOK, "products id", order.Id, "name", order.Name)
		c.Status(http.StatusOK)
		c.JSON(200, orderSuccess(order))
		log.Default()
	}
}
