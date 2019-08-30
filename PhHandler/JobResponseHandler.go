package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMessage"
)

func JobResponseHandler(mh *PhMessage.PhMqttHandler) func(_ interface{}) {
	return func(receive interface{}) {
		//model := receive.(*PhModel.JobResponse)
		//_ = mh.Send(model)
	}
}
