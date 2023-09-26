package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClientMongoDB(ctx context.Context, port string, db string, host string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s", db, host, port)))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func DownMongoDB(ctx context.Context, client *mongo.Client) {
    if err := client.Disconnect(ctx); err != nil {
        panic(err)
    }
}

func NewCollectionMongoDB(client *mongo.Client, nameDb string, nameCollection string) *mongo.Collection {
	collection := client.Database(nameDb).Collection(nameCollection)
	return collection
}
