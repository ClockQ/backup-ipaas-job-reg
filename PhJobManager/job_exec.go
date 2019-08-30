package PhJobManager

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhPanic"
	"os"
	"strconv"
	"strings"
)

func PhJobExec(jobId string,
	kh *PhHelper.PhKafkaHelper,
	mh *PhMqttHelper.PhMqttHelper,
	rh *PhHelper.PhRedisHelper) {

	cStepStr, err := rh.Redis.HGet(jobId, "c_step").Result()
	PhPanic.MqttPanicError(err, mh)
	cStep, err := strconv.Atoi(cStepStr)
	PhPanic.MqttPanicError(err, mh)

	stepStr, err := rh.Redis.HGet(jobId, fmt.Sprintf("step_%d", cStep)).Result()
	PhPanic.MqttPanicError(err, mh)

	process := PhModel.JobProcess{}
	err = json.Unmarshal([]byte(stepStr), &process)
	PhPanic.MqttPanicError(err, mh)

	//err = ProcessExec(jobId, &process, kh)
	PhPanic.MqttPanicError(err, mh)

	err = rh.Redis.HSet(jobId, "c_step", cStep+1).Err()
	PhPanic.MqttPanicError(err, mh)
}

func ProcessExec(jobId string, process *PhModel.JobProcess, kh *PhHelper.PhKafkaHelper) (err error) {
	//jobFunc := GetJobFunc(process.PsType)(kh, process)
	//jobFunc

	switch strings.ToUpper(process.PsType) {
	case "JOB":
		jobRequestTopic := os.Getenv("JOB_REQUEST_TOPIC")
		jobRequest := PhModel.JobRequest{}.GenTMData()
		println(jobRequestTopic)
		println(jobRequest)
		//err = kh.Send(jobRequestTopic, jobRequest)
		if err != nil {
			return
		}
	case "CHANNEL":
		connectRequestTopic := os.Getenv("CONNECT_REQUEST_TOPIC")
		connectRequest := PhModel.ConnectRequest{}.GenTMMongo2HDFSData(jobId)
		println(connectRequestTopic)
		println(connectRequest)
		//err = kh.Send(connectRequestTopic, connectRequest)
		if err != nil {
			return
		}
	}

	return
}
