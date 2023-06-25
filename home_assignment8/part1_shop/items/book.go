package items

import "math/rand"

type Book struct {
	price int
}

func (b Book) GetPrice() int {
	return b.price
}

func (b Book) GetName() string {
	return "book"
}

func NewBook() Book {
	return Book{price: rand.Intn(50) + 1}
}
