# Library API

This library is meant as a relatively simple API for creating a
collection of books and their associated authors.

In researching this, I found
[this as an interesting API](https://www.librarything.com/wiki/index.php/LibraryThing_APIs)
to potentially try and recreate with a hypermedia version.

This will be a "value oriented" API just the sibling Todo library.
What this means, in practice, is:

  * Books and Authors do not know their identity (ID) in the library
  * You cannot modify aspects of the Book or Author.  They are readonly.
  * To change anything about a Book or Author, you must use the Library object.

# Nuances

## Queries

This library API is slightly richer in challenges compared to the Todo
library for a few reasons.  First, it includes multiple domain objects
(book and authors).  Furthermore, there are reasonable queries that
cut across these two domain objects.  For example, I might want to
know all books written by a given author.  This could be a query
against the books to identify those that contain a given author in
their author list or it could be a query against an author for all
books that they have written.  Taken to the extreme, we might want to
find all books whose author's last name is "Smith".  This will
effectively involve a join since we'll first need to find all authors
with the last name Smith (potentially multiple) and then find all
books whose author list includes any of those authors.

## Workflow

Furthermore, the creation of one domain object (books) depends on the
existance of the other domain object (the authors). This creates some
workflow challenges in terms of the API.  It means that creating a
book is at least a two step process.  First, you need to use the API
to find (and if they don't already exist, create) the author and then
you can create the book object.

## Consistency

Operations like Remove should make sure the database is in a
consistent state.  So it should not be allowed to remove an author if
there are books that reference that author.

Similarly, when adding or editing a book, some validation must be done
to make sure that the operation doesn't (re)introduce a dependency on
a non-existent resource.
