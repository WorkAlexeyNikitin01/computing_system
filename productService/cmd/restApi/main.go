package main

import (
	"lab/productService/internal/adapters/postgres"
	"lab/productService/internal/app"
	"log"

	"lab/productService/internal/ports/httpgin"
)

func main() {
	db, err := postgres.NewPostgresDB(postgres.ConfigPostgresProduct{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		log.Fatal(err)
	}
	a := app.NewAppProduct(postgres.NewProductPostgres(db))
	server := httpgin.NewHTTPServer(":18080", a)
	log.Fatal(server.ListenAndServe())
}
