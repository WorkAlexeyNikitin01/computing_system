package postgres

import (
	"lab/orderService/internal/order"

	"github.com/jmoiron/sqlx"
)

type StoreroomPostgres struct {
	db *sqlx.DB
}

func(r *StoreroomPostgres) CreateOrder(products []interface{}) (*order.Order, error) {
	return nil, nil
}
	
func NewStoreroomPostgres(db *sqlx.DB) *StoreroomPostgres {
	return &StoreroomPostgres{
		db: db,
	}
}
