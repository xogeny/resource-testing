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

	// Now remove the database
	err = db.Remove()
	NoError(c, err)
}

func TestInitialize(t *testing.T) {
	Convey("Initialize a TODO database", t, func(c C) {
		db := NewDB("init")

		// Test to make sure we cannot remove an open database
		err := db.Remove()
		IsError(c, err)

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

		_, gstatus, err := db.Get(id1)
		NoError(c, err)
		IsFalse(c, gstatus)

		id2, err := db.AddTodo("Item2", false)
		NoError(c, err)

		_, gstatus, err = db.Get(id2)
		NoError(c, err)
		IsFalse(c, gstatus)

		id3, err := db.AddTodo("Item2", false)
		NoError(c, err)

		_, gstatus, err = db.Get(id3)
		NoError(c, err)
		IsFalse(c, gstatus)

		Equals(c, len(db.ListAll()), 3)
		Equals(c, len(db.ListActive()), 3)
		Equals(c, len(db.ListCompleted()), 0)

		close(db, c)
	})
}
