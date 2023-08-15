package library

import "fmt"

type Shelf interface {
	AddItem(item LibraryItem)
	RemoveItem(item LibraryItem)
}

type BookShelf struct {
	Books []LibraryItem
}

func (bs *BookShelf) AddItem(book LibraryItem) {
	bs.Books = append(bs.Books, book)
}

func (bs *BookShelf) RemoveItem(book LibraryItem) {
	indexToRemove := -1
	for i, b := range bs.Books {
		if b.GetInfo() == book.GetInfo() {
			indexToRemove = i
			break
		}
	}

	if indexToRemove != -1 {
		bs.Books = append(bs.Books[:indexToRemove], bs.Books[indexToRemove+1:]...)
		fmt.Printf("Книгу '%s' видалено з полиці\n", book.GetInfo())
	} else {
		fmt.Printf("Книгу '%s' не знайдено на полиці\n", book.GetInfo())
	}
}
