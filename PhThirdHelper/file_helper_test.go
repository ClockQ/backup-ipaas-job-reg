package PhThirdHelper

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

func TestPhFileHelper(t *testing.T) {
	Convey("测试读取本地文件", t, func() {
		b, err := ioutil.ReadFile("../job_reg.log")
		So(err, ShouldBeNil)

		str := string(b)
		So(str, ShouldNotBeNil)
	})
}
