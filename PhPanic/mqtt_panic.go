package PhPanic

import (
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"log"
	"os"
	"strconv"
)

func MqttPanicError(err error, mh *PhMqttHelper.PhMqttHelper) {
	isTest, _ := strconv.ParseBool(os.Getenv("IS_TEST"))

	if err != nil {
		if isTest { //测试环境直接panic
			panic(err)
		}
		// TODO: 错误协议标准化
		errMsg := fmt.Sprintf("Job Reg 执行出错: %s", err)
		log.Println(errMsg)
		bmlog.StandardLogger().Error(errMsg)
		_ = mh.Send(errMsg)
	}
}
