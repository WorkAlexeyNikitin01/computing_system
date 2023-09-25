package mongodb

import (
	"lab/orderService/internal/order"

	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionMongo struct {
	collection *mongo.Collection
}

func(c *CollectionMongo) CreateOrder(products []interface{}) (*order.Order, error) {
	return nil, nil
}

func NewOrderRepo(collection *mongo.Collection) *CollectionMongo {
	return &CollectionMongo{
		collection: collection,
	}
}
