package mongodb

import (
	"context"
	"lab/orderService/internal/order"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionMongo struct {
	collection *mongo.Collection
}

func(c *CollectionMongo) CreateOrder(order *order.Order) (*mongo.InsertOneResult, error) {
	id, err := c.collection.InsertOne(context.TODO(), bson.D{{Key: "name", Value: order.Name},
															{Key: "quantity", Value: order.Quantity},
															{Key: "price", Value: order.Price},
															{Key: "total", Value: order.Price*float64(order.Quantity)},
															{Key: "code", Value: order.Code},
														})
	if err != nil {
		return nil, err
	}
	log.Println("Success in mongo db")
	return id, nil
}

func NewOrderRepoMongoDb(collection *mongo.Collection) *CollectionMongo {
	return &CollectionMongo{
		collection: collection,
	}
}
