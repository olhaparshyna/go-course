package items

import "math/rand"

type Table struct {
	price int
}

func (t Table) GetPrice() int {
	return t.price
}

func (t Table) GetName() string {
	return "table"
}

func NewTable() Table {
	return Table{price: rand.Intn(100) + 1}
}
