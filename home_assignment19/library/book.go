package library

type LibraryItem interface {
	GetInfo() string
}

type Book struct {
	Title string
}

func (b Book) GetInfo() string {
	return b.Title
}
