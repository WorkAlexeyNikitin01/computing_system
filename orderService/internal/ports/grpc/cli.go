package grpc

import (
	store "lab/storeroomService/grpc"

	"context"
	"log"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Cli store.StoreroomServiceClient
}

func NewClientGRPC() *Client {
	conn, err := ggrpc.DialContext(context.TODO(), "localhost:50054", ggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	cli := store.NewStoreroomServiceClient(conn)
	log.Println("new grpc client ready")
	return &Client{
		Cli: cli,
	}
}