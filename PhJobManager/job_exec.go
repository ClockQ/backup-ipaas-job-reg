package PhJobManager

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhChannel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhPanic"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhThirdHelper"
	"os"
	"strconv"
	"strings"
)

func JobExec(jobId string,
	kh *PhChannel.PhKafkaHelper,
	mh *PhThirdHelper.PhMqttHelper,
	rh *PhThirdHelper.PhRedisHelper) {

	tStepStr, err := rh.Redis.HGet("job_reg_"+jobId, "t_step").Result()
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}
	tStep, err := strconv.Atoi(tStepStr)
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}

	cStepStr, err := rh.Redis.HGet("job_reg_"+jobId, "c_step").Result()
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}
	cStep, err := strconv.Atoi(cStepStr)
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}

	if tStep <= cStep {
		_ = rh.Redis.Del("job_reg_" + jobId).Err()
		_ = mh.Send(PhModel.JobRegResponse{}.SetFinish(jobId, nil))
		return
	}

	stepStr, err := rh.Redis.HGet("job_reg_"+jobId, fmt.Sprintf("step_%d", cStep)).Result()
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}

	process := PhModel.JobProcess{}
	err = json.Unmarshal([]byte(stepStr), &process)
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}

	err = ProcessExec(&process, kh)
	if err != nil {
		PhPanic.MqttPanicError(PhModel.JobRegResponse{}.SetError(jobId, err), mh, rh)
		return
	}
}

func ProcessExec(process *PhModel.JobProcess, kh *PhChannel.PhKafkaHelper) (err error) {
	switch strings.ToUpper(process.PsType) {
	case "CHANNEL":
		connectRequestTopic := os.Getenv("CONNECT_REQUEST_TOPIC")
		connectRequest := PhModel.ConnectRequest{}.New()

		err = connectRequest.Inject(process.Actions)
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

		err = jobRequest.Inject(process.Actions)
		if err != nil {
			return
		}

		err = kh.Send(jobRequestTopic, jobRequest)
		if err != nil {
			return
		}
	case "AGG":
		tmAggRequestTopic := os.Getenv("TMAGG_REQUEST_TOPIC")
		tmAggRequest := PhModel.TmAggRequest{}.New()

		err = tmAggRequest.Inject(process.Actions)
		if err != nil {
			return
		}

		err = kh.Send(tmAggRequestTopic, tmAggRequest)
		if err != nil {
			return
		}
	}

	return
}

func JobExecSuccess(jobId string, rh *PhThirdHelper.PhRedisHelper) (err error) {
	cStepStr, err := rh.Redis.HGet("job_reg_"+jobId, "c_step").Result()
	if err != nil {
		return
	}
	cStep, err := strconv.Atoi(cStepStr)
	if err != nil {
		return
	}
	err = rh.Redis.HSet("job_reg_"+jobId, "c_step", cStep+1).Err()
	if err != nil {
		return
	}
	return
}
