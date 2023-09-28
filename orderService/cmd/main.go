package main

import (
	"context"
	"lab/orderService/internal/adapters/mongodb"
	"lab/orderService/internal/app"
	"log"

	"lab/orderService/internal/ports/grpc"
	"lab/orderService/internal/ports/httpgin"
)

func main() {
	clientMongo, err := mongodb.NewClientMongoDB(context.TODO(), "27017", "mongodb", "localhost")
	if err != nil {
		log.Fatal("error clientMongo")
	}
	db := mongodb.NewCollectionMongoDB(clientMongo, "dbtest", "collectiontest")
	if err != nil {
		log.Fatal(err)
	}
	a := app.NewAppOrder(mongodb.NewOrderRepoMongoDb(db))
	client := grpc.NewClientGRPC()
	server := httpgin.NewHTTPServer(":18081", a, client)
	log.Println("start server")
	log.Fatal(server.ListenAndServe())
}
