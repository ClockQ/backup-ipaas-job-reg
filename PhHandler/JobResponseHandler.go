package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhChannel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhJobManager"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhPanic"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhThirdHelper"
	"strings"
)

func JobResponseHandler(kh *PhChannel.PhKafkaHelper, mh *PhThirdHelper.PhMqttHelper, rh *PhThirdHelper.PhRedisHelper) func(_ interface{}) {
	return func(receive interface{}) {
		model := receive.(*PhModel.JobResponse)
		switch strings.ToUpper(model.Status) {
		case "RUNNING":
			_ = mh.Send(PhModel.JobRegResponse{}.SetRunning(model.JobId, model.Progress, nil))
		case "FINISH":
			final, err := PhJobManager.JobExecSuccess(model.JobId, rh)
			if err != nil {
				PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(model.JobId, err.Error()), mh)
				return
			}
			if final {
				_ = mh.Send(PhModel.JobRegResponse{}.SetFinish(model.JobId, nil))
				return
			}
			go PhJobManager.JobExec(model.JobId, kh, mh, rh)
		case "ERROR":
			PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(model.JobId, model.Message), mh)
			err := PhJobManager.JobExecFatal(model.JobId, rh)
			if err != nil {
				PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(model.JobId, err.Error()), mh)
			}
		default:
			PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(model.JobId, "Job Response 返回状态异常: " + model.Status), mh)
		}
	}
}
