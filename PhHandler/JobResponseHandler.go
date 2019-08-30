package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
)

func JobResponseHandler(mh *PhHelper.PhMqttHelper) func(_ interface{}) {
	return func(receive interface{}) {
		//model := receive.(*PhModel.JobResponse)
		//_ = mh.Send(model)
	}
}
