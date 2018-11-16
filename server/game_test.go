package server

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewGame(t *testing.T) {
	Convey("Given a fresh board", t, func() {
		game:= NewGame()
	})
}
