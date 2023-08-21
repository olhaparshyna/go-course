package repository

import (
	"database/sql"
	"encoding/json"
	"github.com/mediocregopher/radix/v3"
	"log"
)

type ProductRepo interface {
	GetAllProducts() ([]Product, error)
}

type ProductStorage struct {
	*sql.DB
	*radix.Pool
}

func NewProductStorage(db *sql.DB, redis *radix.Pool) *ProductStorage {
	return &ProductStorage{
		DB:   db,
		Pool: redis,
	}
}

type Product struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (storage *ProductStorage) GetAllProducts() ([]Product, error) {
	var products []Product
	var jsonFromRedis []byte

	err := storage.Pool.Do(radix.Cmd(&jsonFromRedis, "GET", "products:all"))

	if err != nil {
		log.Default().Println(err)
	}

	if err == nil && len(jsonFromRedis) > 0 {
		err := json.Unmarshal(jsonFromRedis, &products)
		if err != nil {
			log.Default().Println(err)
		}

		return products, nil
	}

	rows, err := storage.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Println("Error executing SELECT query: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		products = append(products, p)
	}

	json, err := json.Marshal(products)
	if err != nil {
		log.Default().Println(err)
	}

	err = storage.Pool.Do(radix.FlatCmd(nil, "SET", "products:all", json))
	if err != nil {
		log.Default().Println(err)
	}

	err = storage.Pool.Do(radix.FlatCmd(nil, "EXPIRE", "products:all", 3600))
	if err != nil {
		log.Default().Println(err)
	}

	err = storage.Pool.Do(radix.Cmd(&products, "GET", "products:all"))
	if err != nil {
		log.Default().Println(err)
	}

	return products, nil
}
