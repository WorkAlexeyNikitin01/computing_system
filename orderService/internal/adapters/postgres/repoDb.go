package postgres

import (
	"github.com/jmoiron/sqlx"
)

type StoreroomPostgres struct {
	db *sqlx.DB
}
	
func NewStoreroomPostgres(db *sqlx.DB) *StoreroomPostgres {
	return &StoreroomPostgres{
		db: db,
	}
}
