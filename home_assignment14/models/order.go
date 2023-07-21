package models

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

type Store interface {
	Store(Order) error
}

type Processor struct {
	store Store
}

func NewProcessor(s Store) *Processor {
	return &Processor{
		store: s,
	}
}

func (p *Processor) ProcessOrder(o Order) error {
	if err := p.store.Store(o); err != nil {
		return fmt.Errorf("unable to create order: %w", err)
	}

	return nil
}
