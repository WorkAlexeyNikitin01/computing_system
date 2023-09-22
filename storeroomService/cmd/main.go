package main

import (
	"lab/storeroomService/internal/ports/httpgin"
	"lab/storeroomService/internal/adapters/postgres"
	"lab/storeroomService/internal/app"
	"log"
)

func main() {
	dbConfig := postgres.ConfigPostgresStoreroom{
		Host:     "db",
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
	s := httpgin.NewHTTPServer("18081", a)
	log.Fatal(s.ListenAndServe())
}