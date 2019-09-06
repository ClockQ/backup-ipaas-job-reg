package PhPanic

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhThirdHelper"
	"os"
	"strconv"
)

func PanicError(err error, mh *PhThirdHelper.PhMqttHelper) {
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))

	if err != nil {
		if isTest { //测试环境直接panic
			panic(err)
		}
		errMsg := fmt.Sprintf("Job Reg 执行出错: %s", err)
		log.NewLogicLoggerBuilder().Build().Error(errMsg)
		_ = mh.Send(errMsg)
	}
}

func MqttPanicError(model interface{}, mh *PhThirdHelper.PhMqttHelper) {
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))
	if isTest { //测试环境直接panic
		out, _ := json.Marshal(model)
		panic(string(out))
	}

	log.NewLogicLoggerBuilder().Build().Error(model)
	_ = mh.Send(model)
}
