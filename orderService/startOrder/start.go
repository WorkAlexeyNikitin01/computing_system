package startOrder

import (
	"context"
	"lab/notification/rabbitmq/send"
	"lab/orderService/internal/adapters/mongodb"
	"lab/orderService/internal/app"
	"log"

	"lab/orderService/internal/ports/grpc"
	"lab/orderService/internal/ports/httpgin"

	"github.com/rabbitmq/amqp091-go"
)

func Start() {
	log.Println("start order")
	clientMongo, err := mongodb.NewClientMongoDB(context.TODO(), "27017", "mongodb", "mongodb")
	if err != nil {
		log.Fatal("error clientMongo")
	}
	rabbitChan, err := ChanRabbit() 
	if err != nil {
		log.Fatal(err, "error rabbitmq")
	}
	_, _ = send.CreateQueue(rabbitChan)
	db := mongodb.NewCollectionMongoDB(clientMongo, "dbtest", "collectiontest")
	a := app.NewAppOrder(mongodb.NewOrderRepoMongoDb(db))
	client := grpc.NewClientGRPC()
	server := httpgin.NewHTTPServer(":18081", a, client, rabbitChan)
	log.Println("start server")
	log.Fatal(server.ListenAndServe())
}

func ChanRabbit() (*amqp091.Channel, error) {
	conn, err := send.ConnectToRabbit("rabbit")
	if err != nil {
		return nil, err
	}
	
	ch, err := send.CreateChRabbit(conn)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
