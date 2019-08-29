package PhHandler

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMessage"
)

func ConnectResponseHandler(mh *PhMessage.PhMqttHandler) func(_ interface{}) {
	return func(receive interface{}) {
		//model := receive.(*PhModel.ConnectResponse)
		//_ = mh.Send(model)
	}
}
