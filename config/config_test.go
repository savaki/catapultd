package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoad(t *testing.T) {
	Convey("Given a config", t, func() {
		cfg := `{"agent-id":"abc", "auth-token":"123"}`

		Convey("When I call #Load", func() {
			c, err := Load(strings.NewReader(cfg))

			So(err, ShouldBeNil)
			So(c.AgentId, ShouldEqual, "abc")
			So(c.AuthToken, ShouldEqual, "123")
		})
	})
}

func TestLoadFile(t *testing.T) {
	Convey("Given a config file", t, func() {
		cwd, err := filepath.Abs(".")
		So(err, ShouldBeNil)

		filename, err := filepath.Abs(fmt.Sprintf("%s/../examples/catapultd.conf", cwd))
		So(err, ShouldBeNil)

		Convey("When I call #LoadFile", func() {
			c, err := LoadFile(filename)

			So(err, ShouldBeNil)
			So(c.AgentId, ShouldEqual, "abc")
			So(c.AuthToken, ShouldEqual, "123")
		})
	})
}
