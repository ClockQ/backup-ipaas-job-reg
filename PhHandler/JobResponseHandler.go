package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
)

func JobResponseHandler(mh *PhMqttHelper.PhMqttHelper) func(_ interface{}) {
	return func(receive interface{}) {
		//model := receive.(*PhModel.JobResponse)
		//_ = mh.Send(model)
	}
}
