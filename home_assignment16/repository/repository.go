package repository

import (
	"database/sql"
	"github.com/mediocregopher/radix/v3"
)

type Storage struct {
	ProductRepo
	*sql.DB
	*radix.Pool
}

func NewStorage(db *sql.DB, redis *radix.Pool, productStorage ProductRepo) *Storage {
	return &Storage{
		ProductRepo: productStorage,
		DB:          db,
		Pool:        redis,
	}
}
