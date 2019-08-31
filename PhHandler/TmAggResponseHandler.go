package PhHandler

import (
	"errors"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhJobManager"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhPanic"
	"strings"
)

func TmAggResponseHandler(kh *PhHelper.PhKafkaHelper, mh *PhMqttHelper.PhMqttHelper, rh *PhHelper.PhRedisHelper) func(interface{}) {
	return func(receive interface{}) {
		model := receive.(*PhModel.TmAggResponse)
		switch strings.ToUpper(model.Status) {
		case "RUNNING":
			// TODO: 协议标准化
			_ = mh.Send("Agg 执行进度: " + "0")
		case "FINISH":
			err := PhJobManager.JobExecSuccess(model.JobId, rh)
			PhPanic.MqttPanicError(err, mh)
			go PhJobManager.JobExec(model.JobId, kh, mh, rh)
		case "ERROR":
			// TODO: 协议标准化
			PhPanic.MqttPanicError(errors.New("Agg 执行出错: "+model.Message), mh)
		default:
			// TODO: 协议标准化
			PhPanic.MqttPanicError(errors.New("Agg Response 返回状态异常: "+model.Message), mh)
		}
	}
}