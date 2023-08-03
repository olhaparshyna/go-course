package handlers

import (
	"github.com/google/uuid"
	"go-course/home_assignment14/repository"
	"go-course/home_assignment14/requests"
)

func Create(processor repository.Processor, data requests.CreateRequestData) (*repository.Order, error) {
	orderData := repository.Order{
		Id:     uuid.New(),
		UserId: data.UserId,
		Items:  data.Items,
	}

	err := processor.ProcessNewOrder(orderData)

	if err != nil {
		return nil, err
	}

	return &orderData, nil
}
