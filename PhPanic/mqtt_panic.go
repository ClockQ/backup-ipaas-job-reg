package PhPanic

import (
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"log"
	"os"
	"strconv"
)

func MqttPanicError(err error, mh *PhHelper.PhMqttHelper) {
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))

	if err != nil {
		if isTest { //非测试环境才真正发送
			panic(err)
		}
		errMsg := fmt.Sprintf("Job Reg 执行出错: %s", err)
		log.Println(errMsg)
		bmlog.StandardLogger().Error(errMsg)
		_ = mh.Send(errMsg)
	}
}
