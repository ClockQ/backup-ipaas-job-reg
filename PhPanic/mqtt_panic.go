package PhPanic

import (
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"log"
)

func MqttPanicError(err error, mh *PhHelper.PhMqttHelper) {
	if err != nil {
		errMsg := fmt.Sprintf("Job Reg 执行出错: %s", err)
		log.Println(errMsg)
		bmlog.StandardLogger().Error(errMsg)
		_ = mh.Send(errMsg)
	}
}
