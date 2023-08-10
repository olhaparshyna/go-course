package handlers

import (
	"go-course/home_assignment16/repository"
	"log"
)

func ProductList(storage repository.Storage) ([]repository.Product, error) {
	products, err := storage.GetAllProducts()

	if err != nil {
		log.Default().Println(err)
		return nil, err
	}

	return products, nil
}
