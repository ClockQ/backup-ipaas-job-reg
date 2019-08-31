package PhHandler

import (
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
)

func JobResponseHandler(mh *PhMqttHelper.PhMqttHelper) func(_ interface{}) {
	return func(receive interface{}) {
		model := receive.(*PhModel.JobResponse)
		//PhModel.JobResponse{JobId:"cc616902-26ef-7db1-9ada-35bb00e8b727", Status:"Running", Message:"", Progress:"0"}
		fmt.Println(model)
		//_ = mh.Send(model)
	}
}
