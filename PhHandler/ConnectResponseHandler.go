package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
)

func ConnectResponseHandler(mh *PhHelper.PhMqttHelper) func(_ interface{}) {
	return func(receive interface{}) {
		//model := receive.(*PhModel.ConnectResponse)
		//_ = mh.Send(model)
	}
}
