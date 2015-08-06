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

		id, err := db.AddTodo("Initial item to work on", false)
		NoError(c, err)
		NotEquals(c, id, "")

		close(db, c)
	})
}
