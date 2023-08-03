package repository

import (
	"fmt"
	"github.com/google/uuid"
)

type Order struct {
	Id     uuid.UUID `json:"id"`
	UserId int       `json:"userId"`
	Items  []string  `json:"items"`
}

type Orders struct {
	Orders []Order
}

var Storage Orders

func (o *Orders) Store(order Order) error {
	o.Orders = append(o.Orders, order)
	fmt.Println(o.Orders)

	return nil
}
