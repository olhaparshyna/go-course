package library

type LibraryStorage struct {
	books map[string]Book
}

func NewLibrary() *LibraryStorage {
	return &LibraryStorage{
		books: map[string]Book{
			"Маленький принц": Book{Title: "Маленький принц"},
		},
	}
}

func (ls *LibraryStorage) FindBook(title string) *Book {
	book, exists := ls.books[title]
	if !exists {
		return nil
	}
	return &book
}
