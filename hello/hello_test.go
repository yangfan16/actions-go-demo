package hello

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//TestLongestCommonPrefix goconvey单测
func TestLongestCommonPrefix(t *testing.T) {
	Convey("TestLongestCommonPrefix should return fl ", t, func() {
		a := []string{"flower", "flow", "flight"}
		So(LongestCommonPrefix(a), ShouldEqual, "fl")
	})
	Convey("TestLongestCommonPrefix should return empty string ", t, func() {
		b := []string{"dog", "racecar", "car"}
		So(LongestCommonPrefix(b), ShouldEqual, "")
	})
}
