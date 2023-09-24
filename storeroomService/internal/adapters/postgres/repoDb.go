package postgres

import (
	"fmt"
	"lab/storeroomService/internal/storeroom"

	"github.com/jmoiron/sqlx"
)

type StoreroomPostgres struct {
	db *sqlx.DB
}

func(r *StoreroomPostgres) AddToStoreroom(code string, quantity int) (int, error) {
	var id int
	q := fmt.Sprintf("INSERT INTO %s (code, quantity) VALUES ($1, $2) RETURNING id", storeroomTable)
	row := r.db.QueryRow(q, code, quantity)
	if err:=row.Scan(&id); err!=nil {
		return 0, err
	}
	return id, nil
}

func(r *StoreroomPostgres) GetFromStoreroom(code string) (*storeroom.StoreroomProduct, error) {
	var s storeroom.StoreroomProduct
q := fmt.Sprintf("SELECT id, code, quantity FROM %s WHERE code=$1", storeroomTable)
	err := r.db.Get(&s, q, code)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func(r *StoreroomPostgres) DeleteFromStoreroom(code string) (*storeroom.StoreroomProduct, error) {
	var s storeroom.StoreroomProduct
	q := fmt.Sprintf("SELECT id, code, quantity FROM %s WHERE code=$1", storeroomTable)
	err := r.db.Get(&s, q, code)
	if err != nil {
		return nil, err
	}
	q = fmt.Sprintf("DELETE FROM %s WHERE code=$1", storeroomTable)
	_, err = r.db.Exec(q, code)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

	
func NewStoreroomPostgres(db *sqlx.DB) *StoreroomPostgres {
	return &StoreroomPostgres{
		db: db,
	}
}
