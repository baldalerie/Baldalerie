package main

type Library struct {
	Books []Book
}

func (l *Library) GetNames() []string {
	var res []string
	for _, book := range l.Books {
		res = append(res, book.GetName())
	}

	return res
}

type Book struct {
	Name string
}

func (b *Book) GetName() string {
	return b.Name
}

func main() {
	book1 := Book{Name: "Братья Карамазовы"}
	book2 := Book{Name: "Игрок"}

	var books []Book
	books = append(books, book1, book2)

	library := Library{Books: books}

	name := "Братья Карамазовы"

	for _, v := range library.Books {
		if name == v.GetName() {
		}
	}
}
