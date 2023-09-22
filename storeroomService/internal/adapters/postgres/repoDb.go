package postgres

import (

	"github.com/jmoiron/sqlx"
)

type StoreroomPostgres struct {
	db *sqlx.DB
}

func (r *StoreroomPostgres) AddToStoreroom() (interface{}, error) {
	return nil, nil
}

func (r *StoreroomPostgres) GetProductFromStoreroom(code string) (interface{}, error) {
	return nil, nil
}

	
func NewStoreroomPostgres(db *sqlx.DB) *StoreroomPostgres {
	return &StoreroomPostgres{
		db: db,
	}
}
