package postgres

import (
	"lab/storeroomService/internal/storeroom"

	"github.com/jmoiron/sqlx"
)

type StoreroomPostgres struct {
	db *sqlx.DB
}

func(r *StoreroomPostgres) AddToStoreroom(code string, quantity int) (*storeroom.StoreroomProduct, error) {
	return nil, nil
}

func(r *StoreroomPostgres) GetFromStoreroom(code string) (*storeroom.StoreroomProduct, error) {
	return nil, nil
}

func(r *StoreroomPostgres) DeleteFromStoreroom(code string) (*storeroom.StoreroomProduct, error) {
	return nil, nil
}

	
func NewStoreroomPostgres(db *sqlx.DB) *StoreroomPostgres {
	return &StoreroomPostgres{
		db: db,
	}
}
