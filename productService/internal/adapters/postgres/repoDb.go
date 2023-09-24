package postgres

import (
	"fmt"
	"lab/productService/internal/product"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func (r *ProductPostgres) AllListProducts() ([]*product.Product, error) {
	var products []*product.Product
	q := fmt.Sprintf("SELECT * FROM %s", productsTable)
	err := r.db.Get(&products, q)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductPostgres) CreateProduct(p product.Product) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (code, name, price, description) VALUES ($1, $2, $3, $4) RETURNING id", productsTable)
	row := r.db.QueryRow(q, p.Code, p.Name, p.Price, p.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductPostgres) DeleteProduct(code string) (*product.Product, error) {
	var product product.Product
	g := fmt.Sprintf("SELECT id, code, name, price, description FROM %s WHERE code=$1", productsTable)
	err := r.db.Get(&product, g, code) 
	if err != nil {
		return nil, err
	}
	queryDelete := fmt.Sprintf("DELETE FROM %s WHERE code=$1", productsTable)
	_, err = r.db.Exec(queryDelete, code)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{
		db: db,
	}
}
