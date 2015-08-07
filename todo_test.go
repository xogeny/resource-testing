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

		id, err := db.AddTodo(title, false)
		NoError(c, err)
		NotEquals(c, id, "")

		gtitle, gstatus, err := db.Get(id)
		NoError(c, err)
		Equals(c, gtitle, title)
		Equals(c, gstatus, false)

		close(db, c)
	})

	Convey("Check List* methods", t, func(c C) {
		db := NewDB("init")

		id1, err := db.AddTodo("Item1", false)
		NoError(c, err)

		gtitle, gstatus, err := db.Get(id1)
		NoError(c, err)
		Equals(c, gtitle, "Item1")
		IsFalse(c, gstatus)

		id2, err := db.AddTodo("Item2", false)
		NoError(c, err)

		gtitle, gstatus, err = db.Get(id2)
		NoError(c, err)
		Equals(c, gtitle, "Item2")
		IsFalse(c, gstatus)

		id3, err := db.AddTodo("Item3", false)
		NoError(c, err)

		gtitle, gstatus, err = db.Get(id3)
		NoError(c, err)
		Equals(c, gtitle, "Item3")
		IsFalse(c, gstatus)

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

		err = db.Edit(id3, "New Item3", true)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 1)
		Equals(c, len(db.ListCompleted()), 2)

		gtitle, gstatus, err = db.Get(id3)
		NoError(c, err)
		Equals(c, gtitle, "New Item3")
		IsTrue(c, gstatus)

		err = db.Remove(id2)
		NoError(c, err)

		Equals(c, len(db.ListAll()), 2)
		Equals(c, len(db.ListActive()), 0)
		Equals(c, len(db.ListCompleted()), 2)

		// Not idempotent (at least natively)
		err = db.Remove(id2)
		IsError(c, err)

		_, _, err = db.Get(id2)
		IsError(c, err)

		close(db, c)
	})
}
