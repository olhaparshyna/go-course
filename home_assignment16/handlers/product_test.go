package handlers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-course/home_assignment16/repository"
	"testing"
)

type MockProductRepo struct {
	products []repository.Product
	err      error
}

func (m MockProductRepo) GetAllProducts() ([]repository.Product, error) {
	return m.products, m.err
}

func TestProductList_Positive(t *testing.T) {
	mockRepo := MockProductRepo{
		products: []repository.Product{
			{Id: 1, Name: "Product 1", Price: 10},
			{Id: 2, Name: "Product 2", Price: 20},
		},
		err: nil,
	}

	products, err := ProductList(mockRepo)
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 2", products[1].Name)
}

func TestProductList_Negative(t *testing.T) {
	mockRepo := MockProductRepo{
		products: nil,
		err:      errors.New("failed to get products"),
	}

	products, err := ProductList(mockRepo)
	assert.Error(t, err)
	assert.Nil(t, products)
}
