package postgres

import (
	"fmt"
	"lab/first_microservice/internal/product"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func (r *ProductPostgres) AllListProducts() ([]product.Product, error) {
	var products []product.Product
	q := fmt.Sprintf("SELECT * FROM %s", productsTable)
	err := r.db.Get(&products, q)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewAuthPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{
		db: db,
	}
}
