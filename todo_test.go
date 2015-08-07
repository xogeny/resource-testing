package gotodo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/xogeny/xconvey"
)

func close(db *TodoDatabase, c C) {
	// Make sure we can close the database
	err := db.Close()
	NoError(c, err)
}

func TestInitialize(t *testing.T) {
	Convey("Initialize a TODO database", t, func(c C) {
		db := NewDB("init")

		close(db, c)
	})
}

func TestPopulate(t *testing.T) {
	Convey("Populate a TODO database", t, func(c C) {
		db := NewDB("init")
		title := "Initial item to work on"

		id, err := db.AddTodo(MakeItem(title, false))
		NoError(c, err)
		NotEquals(c, id, "")

		item, err := db.Get(id)
		NoError(c, err)
		Equals(c, item.Title(), title)
		Equals(c, item.Completed(), false)

		close(db, c)
	})

	Convey("Check List* methods", t, func(c C) {
		db := NewDB("init")

		id1, err := db.AddTodo(MakeItem("Item1", false))
		NoError(c, err)

		item, err := db.Get(id1)
		NoError(c, err)
		Equals(c, item.Title(), "Item1")
		IsFalse(c, item.Completed())

		id2, err := db.AddTodo(MakeItem("Item2", false))
		NoError(c, err)

		item, err = db.Get(id2)
		NoError(c, err)
		Equals(c, item.Title(), "Item2")
		IsFalse(c, item.Completed())

		id3, err := db.AddTodo(MakeItem("Item3", false))
		NoError(c, err)

		item, err = db.Get(id3)
		NoError(c, err)
		Equals(c, item.Title(), "Item3")
		IsFalse(c, item.Completed())

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 3)
		Equals(c, len(db.ListCompleted()), 0)

		err = db.MarkCompleted(id1)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 2)
		Equals(c, len(db.ListCompleted()), 1)

		err = db.MarkCompleted(id3)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 1)
		Equals(c, len(db.ListCompleted()), 2)

		// Test idempotency
		err = db.MarkCompleted(id3)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 1)
		Equals(c, len(db.ListCompleted()), 2)

		err = db.ClearCompleted(id3)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 2)
		Equals(c, len(db.ListCompleted()), 1)

		err = db.Edit(id3, MakeItem("New Item3", true))
		NoError(c, err)

		item, err = db.Get(id3)
		NoError(c, err)
		Equals(c, item.Title(), "New Item3")
		IsTrue(c, item.Completed())

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 1)
		Equals(c, len(db.ListCompleted()), 2)

		item, err = db.Get(id3)
		NoError(c, err)
		Equals(c, item.Title(), "New Item3")
		IsTrue(c, item.Completed())

		err = db.Remove(id2)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 2)
		Equals(c, len(db.ListActive()), 0)
		Equals(c, len(db.ListCompleted()), 2)

		// Not idempotent (at least natively)
		err = db.Remove(id2)
		IsError(c, err)

		_, err = db.Get(id2)
		IsError(c, err)

		// Now we test operations on non-existent ids to get
		// better coverage
		bogus := ID("Can't be an id")

		err = db.Edit(bogus, MakeItem("Bogus Item", true))
		IsError(c, err)

		err = db.MarkCompleted(bogus)
		IsError(c, err)

		err = db.ClearCompleted(bogus)
		IsError(c, err)

		close(db, c)
	})
}
