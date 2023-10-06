package httpgin

import (
	"context"
	"encoding/json"
	"lab/notification/rabbitmq/send"
	"lab/orderService/internal/app"
	"lab/orderService/internal/order"

	store "lab/storeroomService/grpc"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func createOrder(a app.AppOrderInterface, client store.StoreroomServiceClient, rabbitChan *amqp091.Channel) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody OrderRequest
		err := c.Bind(&reqBody)
		if err != nil {
			log.Println("error create order", err)
			c.JSON(400, orderError(err))
			return
		}

		sproduct, _ := client.GetFromStoreroom(context.TODO(), &store.StoreProductRequest{Code: reqBody.Code})
		if reqBody.Quantity > int(sproduct.Quantity) {
			log.Println("not product in store", err)
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
		body := map[string]any{
			"id": order.Id,
			"quantity":   order.Quantity,
			"code":    order.Code,
			"name": order.Name,
			"price": order.Price,
		}
		byteBody, _ := json.Marshal(body)
		_ = send.SendToQueue(byteBody, context.TODO(), rabbitChan)
		log.Println("Success create order", http.StatusOK, "products id", order.Id, "name", order.Name)
		c.Status(http.StatusOK)
		c.JSON(200, orderSuccess(order))
		log.Default()
	}
}
