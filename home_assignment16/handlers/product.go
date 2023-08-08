package handlers

import "go-course/home_assignment16/repository"

func ProductList(storage repository.Storage) ([]repository.Product, error) {
	products, err := storage.GetAllProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
