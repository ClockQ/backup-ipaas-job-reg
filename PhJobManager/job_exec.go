package PhJobManager

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhPanic"
	"github.com/mitchellh/mapstructure"
	"os"
	"strconv"
	"strings"
)

func PhJobExec(jobId string,
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

	err = ProcessExec(jobId, &process, kh)
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}

	// todo 回调时执行
	err = rh.Redis.HSet(jobId, "c_step", cStep+1).Err()
	if err != nil {
		PhPanic.MqttPanicError(err, mh)
		return
	}

	if tStep == cStep+1 {
		_ = mh.Send(fmt.Sprintf("%s 执行完成", jobId))
	}
}

func ProcessExec(jobId string, process *PhModel.JobProcess, kh *PhHelper.PhKafkaHelper) (err error) {
	switch strings.ToUpper(process.PsType) {
	case "CHANNEL":
		connectRequestTopic := os.Getenv("CONNECT_REQUEST_TOPIC")
		connectRequest := PhModel.ConnectRequest{}.New()

		err = mapstructure.Decode(process.JobConfig, connectRequest)
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

		err = mapstructure.Decode(process.JobConfig, jobRequest)
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
