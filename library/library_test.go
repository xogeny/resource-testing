package library

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/xogeny/xconvey"
)

func TestLibrary(t *testing.T) {
	Convey("Testing library API", t, func(c C) {
		bogusAuthor := AuthorID("invalid author id")
		bogusBook := BookID("invalid book id")

		lib := NewLibrary()
		NotNil(c, lib)

		// -- Test Author API --
		mmt, err := lib.AddAuthor(MakeAuthor("Michael", "Tiller"))
		NoError(c, err)
		me, err := lib.GetAuthor(mmt)
		NoError(c, err)
		Equals(c, me.FullName(), "Michael Tiller")

		_, err = lib.GetAuthor(bogusAuthor)
		IsError(c, err)

		err = lib.EditAuthor(mmt, MakeAuthor("Michael", "Tiller", "Mark"))
		NoError(c, err)
		me, err = lib.GetAuthor(mmt)
		NoError(c, err)
		Equals(c, me.FullName(), "Michael Mark Tiller")

		err = lib.EditAuthor(bogusAuthor, MakeAuthor("Bogus", "Author"))
		IsError(c, err)

		// -- Test Book API --
		mbeid, err := lib.AddBook(MakeBook("Modelica by Example", "abc123", mmt))
		NoError(c, err)

		_, err = lib.AddBook(MakeBook("Modelica by Example", "def897", bogusAuthor))
		IsError(c, err)

		mbe, err := lib.GetBook(mbeid)
		NoError(c, err)
		Equals(c, mbe.Title(), "Modelica by Example")
		Equals(c, mbe.ISBN(), "abc123")
		Resembles(c, mbe.Authors(), []AuthorID{mmt})
		Equals(c, len(mbe.Authors()), 1)

		_, err = lib.GetBook(bogusBook)
		IsError(c, err)

		err = lib.EditBook(mbeid, MakeBook("Modelica by Example", "abc789", bogusAuthor))
		IsError(c, err)

		err = lib.EditBook(mbeid, MakeBook("Modelica by Example", "xyz123", mmt))
		NoError(c, err)
		mbe, err = lib.GetBook(mbeid)
		NoError(c, err)
		Equals(c, mbe.Title(), "Modelica by Example")
		Equals(c, mbe.ISBN(), "xyz123")
		Resembles(c, mbe.Authors(), []AuthorID{mmt})
		Equals(c, len(mbe.Authors()), 1)

		err = lib.EditBook(bogusBook, mbe)
		IsError(c, err)

		// -- Test List API --
		Equals(c, len(lib.ListAuthors()), 1)
		Resembles(c, lib.ListAuthors(), []AuthorID{mmt})

		Equals(c, len(lib.ListAuthorsByName("Michael")), 1)
		Resembles(c, lib.ListAuthorsByName("Michael"), []AuthorID{mmt})
		Equals(c, len(lib.ListAuthorsByName("Mark")), 1)
		Resembles(c, lib.ListAuthorsByName("Mark"), []AuthorID{mmt})
		Equals(c, len(lib.ListAuthorsByName("Tiller")), 1)
		Resembles(c, lib.ListAuthorsByName("Tiller"), []AuthorID{mmt})
		Equals(c, len(lib.ListAuthorsByName("tiller")), 1)
		Resembles(c, lib.ListAuthorsByName("tiller"), []AuthorID{mmt})
		Equals(c, len(lib.ListAuthorsByName("Fitzgerald")), 0)
		Resembles(c, lib.ListAuthorsByName("Fitzgerald"), []AuthorID{})

		Equals(c, len(lib.ListBooks()), 1)
		Resembles(c, lib.ListBooks(), []BookID{mbeid})

		Equals(c, len(lib.ListBooksByAuthor(mmt)), 1)
		Resembles(c, lib.ListBooksByAuthor(mmt), []BookID{mbeid})
		Equals(c, len(lib.ListBooksByAuthor(bogusAuthor)), 0)
		Resembles(c, lib.ListBooksByAuthor(bogusAuthor), []BookID{})

		Equals(c, len(lib.ListBooksByAuthorName("Tiller")), 1)
		Resembles(c, lib.ListBooksByAuthorName("Tiller"), []BookID{mbeid})

		Equals(c, len(lib.ListBooksByTitle("Modelica by Example", true)), 1)
		Resembles(c, lib.ListBooksByTitle("Modelica by Example", true), []BookID{mbeid})
		Equals(c, len(lib.ListBooksByTitle("Modelica", true)), 0)
		Resembles(c, lib.ListBooksByTitle("Modelica", true), []BookID{})
		Equals(c, len(lib.ListBooksByTitle("modelica", false)), 1)
		Resembles(c, lib.ListBooksByTitle("modelica", false), []BookID{mbeid})

		Equals(c, len(lib.ListBooksByISBN("xyz123")), 1)
		Resembles(c, lib.ListBooksByISBN("xyz123"), []BookID{mbeid})
		Equals(c, len(lib.ListBooksByISBN("abc123")), 0)
		Resembles(c, lib.ListBooksByISBN("abc123"), []BookID{})

		// -- Test Remove methods --
		// Can't remove because it is referenced by an existing book
		err = lib.RemoveAuthor(mmt)
		IsError(c, err)

		err = lib.RemoveBook(mbeid)
		NoError(c, err)

		err = lib.RemoveBook(bogusBook)
		IsError(c, err)

		err = lib.RemoveBook(mbeid)
		IsError(c, err)

		err = lib.RemoveAuthor(mmt)
		NoError(c, err)

		err = lib.RemoveAuthor(bogusAuthor)
		IsError(c, err)

		err = lib.RemoveAuthor(mmt)
		IsError(c, err)
	})
}
