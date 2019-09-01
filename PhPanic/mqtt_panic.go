package PhPanic

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhThirdHelper"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"log"
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
		log.Println(errMsg)
		bmlog.StandardLogger().Error(errMsg)
		_ = mh.Send(errMsg)
	}
}

func MqttPanicError(model interface{}, mh *PhThirdHelper.PhMqttHelper) {
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))
	if isTest { //测试环境直接panic
		out, _ := json.Marshal(model)
		panic(string(out))
	}

	log.Println(model)
	bmlog.StandardLogger().Error(model)
	_ = mh.Send(model)
}
