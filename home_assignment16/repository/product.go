package repository

import (
	"log"
)

type Product struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (storage *Storage) GetAllProducts() ([]Product, error) {
	rows, err := storage.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Println("Error executing SELECT query: ", err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
