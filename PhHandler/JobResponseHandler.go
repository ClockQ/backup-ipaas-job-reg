package PhHandler

import (
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
)

func JobResponseHandler(mh *PhMqttHelper.PhMqttHelper) func(_ interface{}) {
	return func(receive interface{}) {
		model := receive.(*PhModel.JobResponse)
		fmt.Println(model)
		//_ = mh.Send(model)
	}
}
