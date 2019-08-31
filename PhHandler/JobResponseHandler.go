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

func JobResponseHandler(kh *PhHelper.PhKafkaHelper, mh *PhMqttHelper.PhMqttHelper, rh *PhHelper.PhRedisHelper) func(_ interface{}) {
	return func(receive interface{}) {
		model := receive.(*PhModel.JobResponse)
		switch strings.ToUpper(model.Status) {
		case "RUNNING":
			_ : mh.Send("Job 执行进度: " + model.Progress)
		case "FINISH":
			err := PhJobManager.JobExecSuccess(model.JobId, rh)
			PhPanic.MqttPanicError(err, mh)
			go PhJobManager.JobExec(model.JobId, kh, mh, rh)
		case "ERROR":
			PhPanic.MqttPanicError(errors.New("Job 执行出错: " + model.Message), mh)
		default:
			PhPanic.MqttPanicError(errors.New("Job Response 返回状态异常: " + model.Message), mh)
		}
	}
}
