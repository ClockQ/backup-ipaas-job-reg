package PhPanic

import (
	"errors"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestPanicError(t *testing.T) {
	_ = os.Setenv("IS_TEST", "true")

	Convey("测试 PanicError", t, func() {
		PanicError(errors.New("测试 PanicError"), nil)
	})
}

func TestMqttPanicError(t *testing.T) {
	_ = os.Setenv("IS_TEST", "true")

	Convey("测试 MqttPanicError", t, func() {
		model := PhModel.JobRegResponse{}.SetError("JobId", "测试 MqttPanicError")
		MqttPanicError(model, nil, nil)
	})
}
