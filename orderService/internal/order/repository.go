package order

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	CreateOrder(order *Order) (*mongo.InsertOneResult, error)
}
