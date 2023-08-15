package main

import (
	"fmt"
	"go-course/home_assignment19/library"
)

func main() {
	// Створення бази даних з книгами
	db := library.NewLibrary()

	// Створення менеджера бібліотеки
	librarian := &library.Manager{
		Database: db,
		Shelf:    &library.BookShelf{},
	}

	// Створення клієнтів
	client1 := library.Client{Name: "Клієнт 1"}
	client2 := library.Client{Name: "Клієнт 2"}

	book1 := librarian.LendBook("Маленький принц")
	book2 := library.Book{Title: "Unknown"}

	err := librarian.ReturnBook(client2, book2)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = librarian.ReturnBook(client1, book1)

	if err != nil {
		fmt.Println(err.Error())
	}
}
