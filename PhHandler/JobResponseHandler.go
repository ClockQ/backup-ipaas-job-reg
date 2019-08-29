package PhHandler

import (
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
)

func JobResponseHandler(receive interface{}) {
	model := receive.(*PhModel.JobResponse)
	fmt.Printf("receive msg : %#v \n", model)
}
