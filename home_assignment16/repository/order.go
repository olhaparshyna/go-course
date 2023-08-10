package repository

import (
	"github.com/google/uuid"
	"log"
	"strings"
)

type Order struct {
	Id     uuid.UUID `json:"id"`
	UserId int64     `json:"userId"`
	Items  []string  `json:"items"`
}

func (storage *Storage) OrderStore(id int64, items []string) (error, *Order) {
	order := Order{
		Id:     uuid.New(),
		UserId: id,
		Items:  items,
	}
	insertQuery := "INSERT INTO orders (id, userId, items) VALUES (?, ?, ?)"

	result, err := storage.DB.Exec(insertQuery, order.Id.String(), order.UserId, strings.Join(order.Items, ","))
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	_, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err, nil
	}
	return err, &order
}
