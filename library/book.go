package library

type Book struct {
	title   string
	isbn    string
	authors []AuthorID
}

func (b Book) Title() string {
	return b.title
}

func (b Book) ISBN() string {
	return b.isbn
}

func (b Book) Authors() []AuthorID {
	return b.authors
}

func MakeBook(title string, isbn string, authors ...AuthorID) Book {
	return Book{
		title:   title,
		isbn:    isbn,
		authors: authors,
	}
}
