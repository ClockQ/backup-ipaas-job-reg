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

// TODO Delete
func MqttPanicError_del(err error, mh *PhThirdHelper.PhMqttHelper) {
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

func MqttPanicError(model interface{}, mh *PhThirdHelper.PhMqttHelper, rh *PhThirdHelper.PhRedisHelper) {
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))
	if isTest { //测试环境直接panic
		out, _ := json.Marshal(model)
		panic(string(out))
	}

	//TODO: 提取本次 Job 的执行策略，保留现场（暂无）
	log.Println(model)
	bmlog.StandardLogger().Error(model)
	_ = mh.Send(model)
}
