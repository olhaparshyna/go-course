package repository

import (
	"fmt"
)

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

func (p *Processor) ProcessNewOrder(o Order) error {
	if err := p.store.Store(o); err != nil {
		return fmt.Errorf("unable to create order: %w", err)
	}

	return nil
}
