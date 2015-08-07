package library

import (
	"fmt"
	uuid4 "github.com/nathanwinther/go-uuid4"
	"strings"
)

// This is here just to summarize the interface for a Library
type LibraryInterface interface {
	// Author related API
	GetAuthor(AuthorID) (Author, error)
	AddAuthor(Author) (AuthorID, error)
	EditAuthor(AuthorID, Author) error
	RemoveAuthor(AuthorID) error

	// Book related API
	GetBook(BookID) (Book, error)
	AddBook(Book) (BookID, error)
	EditBook(BookID, Book) error
	RemoveBook(BookID) error

	// Queries
	ListAuthors() []AuthorID
	ListAuthorsByName(string) []AuthorID

	ListBooks() []BookID
	ListBooksByAuthor(AuthorID) []BookID
	ListBooksByAuthorName(name string) []BookID
	ListBooksByTitle(title string, exact bool) []BookID
	ListBooksByISBN(isbn string) []BookID
}

// Handle to identify a book
type BookID string

// Handle to identify an author
type AuthorID string

// This is the actual library type used to manage
// books and authors
type Library struct {
	authors map[AuthorID]Author
	books   map[BookID]Book
}

func (l *Library) GetAuthor(id AuthorID) (Author, error) {
	author, exists := l.authors[id]
	if !exists {
		return author, fmt.Errorf("No author associated with id %v", id)
	}
	return author, nil
}

func (l *Library) AddAuthor(a Author) (AuthorID, error) {
	uuid, err := uuid4.New()
	if err != nil {
		return AuthorID(""), err
	}
	id := AuthorID(uuid)
	l.authors[id] = a
	return id, nil
}

func (l *Library) EditAuthor(id AuthorID, author Author) error {
	_, exists := l.authors[id]
	if !exists {
		return fmt.Errorf("No author associated with id %v", id)
	}
	l.authors[id] = author
	return nil
}

func (l *Library) RemoveAuthor(id AuthorID) error {
	_, exists := l.authors[id]
	if !exists {
		return fmt.Errorf("No author associated with id %v", id)
	}
	for _, b := range l.books {
		if b.HasAuthor(id) {
			return fmt.Errorf("Cannot remove author because book %v depends on it", b)
		}
	}
	delete(l.authors, id)
	return nil
}

func (l *Library) GetBook(id BookID) (Book, error) {
	book, exists := l.books[id]
	if !exists {
		return book, fmt.Errorf("No book associated with id %v", id)
	}
	return book, nil
}

func (l *Library) AddBook(b Book) (BookID, error) {
	uuid, err := uuid4.New()
	if err != nil {
		return BookID(""), err
	}

	for _, ba := range b.authors {
		if _, exists := l.authors[ba]; !exists {
			return BookID(""),
				fmt.Errorf("Cannot add book %s to include non-existent author %v",
					b.title, ba)
		}
	}

	id := BookID(uuid)
	l.books[id] = b
	return id, nil
}

func (l *Library) EditBook(id BookID, book Book) error {
	_, exists := l.books[id]
	if !exists {
		return fmt.Errorf("No books associated with id %v", id)
	}
	for _, ba := range book.authors {
		if _, exists := l.authors[ba]; !exists {
			return fmt.Errorf("Cannot edit book %v to include non-existent author %v",
				id, ba)
		}
	}
	l.books[id] = book
	return nil
}

func (l *Library) RemoveBook(id BookID) error {
	_, exists := l.books[id]
	if !exists {
		return fmt.Errorf("No book associated with id %v", id)
	}
	delete(l.books, id)
	return nil
}

func (l *Library) ListAuthors() []AuthorID {
	ret := []AuthorID{}
	for aid, _ := range l.authors {
		ret = append(ret, aid)
	}
	return ret
}

func (l *Library) ListAuthorsByName(name string) []AuthorID {
	ret := []AuthorID{}
	for aid, author := range l.authors {
		if author.Matches(name) {
			ret = append(ret, aid)
		}
	}
	return ret
}

type bookPredicate func(book Book) bool

func (l *Library) listBooks(f bookPredicate) []BookID {
	ret := []BookID{}
	for bid, book := range l.books {
		if f(book) {
			ret = append(ret, bid)
		}
	}
	return ret
}

func (l *Library) ListBooks() []BookID {
	return l.listBooks(func(book Book) bool { return true })
}

func (l *Library) ListBooksByAuthor(id AuthorID) []BookID {
	return l.listBooks(func(book Book) bool { return book.HasAuthor(id) })
}

func (l *Library) ListBooksByAuthorName(name string) []BookID {
	authors := l.ListAuthorsByName(name)
	return l.listBooks(func(book Book) bool {
		return book.HasAuthor(authors...)
	})
}

func (l *Library) ListBooksByTitle(title string, exact bool) []BookID {
	ltitle := strings.ToLower(title)
	if exact {
		return l.listBooks(func(book Book) bool {
			return strings.ToLower(book.Title()) == ltitle
		})
	} else {
		return l.listBooks(func(book Book) bool {
			btitle := strings.ToLower(book.Title())

			return strings.Contains(btitle, ltitle)
		})
	}
}

func (l *Library) ListBooksByISBN(isbn string) []BookID {
	return l.listBooks(func(book Book) bool {
		return strings.ToLower(book.ISBN()) == strings.ToLower(isbn)
	})
}

func NewLibrary() *Library {
	ret := Library{
		authors: map[AuthorID]Author{},
		books:   map[BookID]Book{},
	}
	return &ret
}

var _ LibraryInterface = (*Library)(nil)
