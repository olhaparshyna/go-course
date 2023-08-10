package repository

import (
	"database/sql"
	"github.com/mediocregopher/radix/v3"
)

type Repository interface {
	GetAllProducts() ([]Product, error)
	FindUserIdByEmail(email string) int64
	UserStore(name string, email string) (error, int64)
	OrderStore(userId int64, items []string) (error, *Order)
}

type Storage struct {
	*sql.DB
	*radix.Pool
}

func NewStorage(db *sql.DB, redis *radix.Pool) *Storage {
	return &Storage{
		DB:   db,
		Pool: redis,
	}
}
