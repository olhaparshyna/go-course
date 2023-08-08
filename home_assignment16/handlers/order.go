package handlers

import (
	"go-course/home_assignment16/repository"
	"go-course/home_assignment16/requests"
)

func Create(storage repository.Storage, data requests.CreateRequestData) (*repository.Order, error) {
	userId := storage.FindUserIdByEmail(data.Email)
	if userId == 0 {
		err, newUserId := storage.UserStore(data.Name, data.Email)
		if err != nil {
			return nil, err
		}

		userId = newUserId
	}

	err, order := storage.OrderStore(userId, data.Email, data.Items)

	if err != nil {
		return nil, err
	}

	return order, nil
}
