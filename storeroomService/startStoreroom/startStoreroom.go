package startStoreroom

import (
	"lab/storeroomService/internal/adapters/postgres"
	"lab/storeroomService/internal/app"
	gport "lab/storeroomService/grpc"
	"lab/storeroomService/internal/ports/httpgin"
	
	"net"
	"log"

	"google.golang.org/grpc"
)

func Start() {
	log.Println("start storeroom")
	dbConfig := postgres.ConfigPostgresStoreroom{
		Host:     "postgres-storeroom",
		Port:     "5433",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	}
	pg, err := postgres.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	a := app.NewAppStoreroom(postgres.NewStoreroomPostgres(pg))
	s := httpgin.NewHTTPServer(":18082", a)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	svc := gport.NewService(a)
	gport.RegisterStoreroomServiceServer(grpcServer, svc)

	go func() {
		log.Println("start grpc store port 50054")
		log.Fatal(grpcServer.Serve(lis))
	}()

	log.Println("start service store")
	log.Fatal(s.ListenAndServe())
}
