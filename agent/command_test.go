package agent

import (
	"strings"

	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestContext(t *testing.T) {
	Convey("Given a Context", t, func() {
		c := Context{
			Dir: "../examples",
			Log: StdoutLogFunc,
		}

		Convey("When I call #Command on a valid file", func() {
			command := "echo.sh"
			cmd, err := c.Command(command)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
				So(cmd, ShouldNotBeNil)
			})

			Convey("And I expect that executing the command should return 'hello world'", func() {
				output, err := cmd.CombinedOutput()

				Convey("Then I expect no errors", func() {
					So(err, ShouldBeNil)
				})

				Convey("And I expect hello world", func() {
					text := strings.TrimSpace(string(output))
					So(text, ShouldEqual, "hello world")
				})
			})
		})
	})
}
