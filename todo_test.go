package gotodo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/xogeny/xconvey"
)

func TestInitialize(t *testing.T) {
	Convey("Initialize a TODO database", t, func(c C) {
		db := NewDB("init")
		Equals(c, 5, 5)
		err := db.Close()
		NoError(c, err)
		err = db.Remove()
		NoError(c, err)
	})
}
