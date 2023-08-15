package library

import (
	"errors"
	"fmt"
)

type Manager struct {
	Database Library
	Shelf    Shelf
}

func (lm *Manager) LendBook(title string) LibraryItem {
	book := lm.Database.FindBook(title)
	lm.Shelf.AddItem(*book)
	lm.Shelf.RemoveItem(book)

	return book
}

func (lm *Manager) ReturnBook(client Client, book LibraryItem) error {
	foundBook := lm.Database.FindBook(book.GetInfo())
	if foundBook == nil {
		return errors.New("Книга не зараєстрована у бібліотеці")
	}

	lm.Shelf.AddItem(book)
	fmt.Printf("%s повернув книгу '%s'\n", client.Name, book.GetInfo())
	return nil
}
