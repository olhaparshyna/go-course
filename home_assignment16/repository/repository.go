package repository

import "database/sql"

type Repository interface {
	GetAllProducts() ([]Product, error)
	FindUserIdByEmail(email string) int64
	UserStore(name string, email string) (error, int64)
	OrderStore(userId *int64, email string, items []string) (error, int64)
}

type Storage struct {
	*sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}
