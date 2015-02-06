package agent

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewAgent(t *testing.T) {
	Convey("Given a config", t, func() {
		So(1, ShouldEqual, 1)
	})
}
