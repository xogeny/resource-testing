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

// Returns true if any of the authors in the list 'ids' is an author
// of this book.
func (b Book) HasAuthor(ids ...AuthorID) bool {
	for _, id := range ids {
		for _, aid := range b.authors {
			if aid == id {
				return true
			}
		}
	}
	return false
}

func MakeBook(title string, isbn string, authors ...AuthorID) Book {
	return Book{
		title:   title,
		isbn:    isbn,
		authors: authors,
	}
}
