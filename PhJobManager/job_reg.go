package PhJobManager

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"time"
)

func PhJobReg(model PhModel.JobReg,
	_ *PhHelper.PhKafkaHelper,
	_ *PhHelper.PhMqttHelper,
	rh *PhHelper.PhRedisHelper) (err error) {

	jobId := model.JobId

	err = rh.Redis.HSet(jobId, "c_step", 0).Err()
	if err != nil {
		return
	}

	err = rh.Redis.HSet(jobId, "t_step", len(model.Process)).Err()
	if err != nil {
		return
	}

	for i, v := range model.Process {
		value, err := json.Marshal(v)
		if err != nil {
			return err
		}

		err = rh.Redis.HSet(jobId, fmt.Sprintf("step_%d", i), value).Err()
		if err != nil {
			return err
		}
	}

	// job注册信息 24h 过期
	err = rh.Redis.Expire(jobId, time.Hour*24).Err()
	if err != nil {
		return
	}
	return
}
