package library

// This is here just to summarize the interface for a Library
type LibraryInterface interface {
}

// Handle to identify a book
type BookID string

// Handle to identify an author
type AuthorID string

// This is the actual library type used to manage
// books and authors
type Library struct {
}

var _ LibraryInterface = (*Library)(nil)
