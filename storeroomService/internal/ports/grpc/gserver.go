package grpc

import (
	"context"
	"lab/storeroomService/internal/app"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type gRPCServer interface {
	StoreroomServiceServer
}

type gRPCServerStruct struct {
	A app.AppStoreroomInterface
	UnimplementedStoreroomServiceServer
}

func (g *gRPCServerStruct) GetFromStoreroom(ctx context.Context, req *StoreProductRequest) (*SProductResponse, error) {
	sproduct, err := g.A.GetFromStoreroom(req.Code)
	if err != nil {
		log.Println("error in create ad ", err)
		return nil, status.Error(codes.InvalidArgument, "error create ad")
	}
	log.Printf("user %v && create ad %v \n", sproduct.Id, sproduct.Code)
	return &SProductResponse{Id: int64(sproduct.Id), Code: sproduct.Code, Quantity: int64(sproduct.Quantity)}, nil
}

func NewService(a app.AppStoreroomInterface) gRPCServer {
	return &gRPCServerStruct{A: a}
}
