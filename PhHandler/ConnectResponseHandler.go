package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
)

func ConnectResponseHandler(mh *PhMqttHelper.PhMqttHelper) func(_ interface{}) {
	return func(receive interface{}) {
		//model := receive.(*PhModel.ConnectResponse)
		//_ = mh.Send(model)
	}
}
