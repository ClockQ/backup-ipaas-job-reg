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

func JobExec(jobId string,
	kh *PhHelper.PhKafkaHelper,
	mh *PhMqttHelper.PhMqttHelper,
	rh *PhHelper.PhRedisHelper) {

	tStepStr, err := rh.Redis.HGet(jobId, "t_step").Result()
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}
	tStep, err := strconv.Atoi(tStepStr)
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}

	cStepStr, err := rh.Redis.HGet(jobId, "c_step").Result()
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}
	cStep, err := strconv.Atoi(cStepStr)
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}

	if tStep == cStep {
		_ = mh.Send(fmt.Sprintf("%s 执行完成", jobId))
		return
	}

	stepStr, err := rh.Redis.HGet(jobId, fmt.Sprintf("step_%d", cStep)).Result()
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}

	process := PhModel.JobProcess{}
	err = json.Unmarshal([]byte(stepStr), &process)
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}

	err = ProcessExec(&process, kh)
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}
}

func ProcessExec(process *PhModel.JobProcess, kh *PhHelper.PhKafkaHelper) (err error) {
	switch strings.ToUpper(process.PsType) {
	case "CHANNEL":
		connectRequestTopic := os.Getenv("CONNECT_REQUEST_TOPIC")
		connectRequest := PhModel.ConnectRequest{}.New()

		err = connectRequest.Inject(process.JobConfig)
		if err != nil {
			return
		}

		err = kh.Send(connectRequestTopic, connectRequest)
		if err != nil {
			return
		}
	case "JOB":
		jobRequestTopic := os.Getenv("JOB_REQUEST_TOPIC")
		jobRequest := PhModel.JobRequest{}.New()

		err = jobRequest.Inject(process.JobConfig)
		if err != nil {
			return
		}

		err = kh.Send(jobRequestTopic, jobRequest)
		if err != nil {
			return
		}
	}

	return
}

func ProcessStatus() (err error) {
	return nil
}
