package library

type Book struct {
	Title string
}

func (b Book) GetInfo() string {
	return b.Title
}
