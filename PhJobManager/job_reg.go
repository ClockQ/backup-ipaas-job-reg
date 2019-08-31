package PhJobManager

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"time"
)

func PhJobReg(model PhModel.JobReg,
	_ *PhHelper.PhKafkaHelper,
	_ *PhMqttHelper.PhMqttHelper,
	rh *PhHelper.PhRedisHelper) (err error) {

	jobId := model.JobId

	_ = rh.Redis.Del("job_reg_"+jobId)

	err = rh.Redis.HSet("job_reg_"+jobId, "c_step", 0).Err()
	if err != nil {
		return
	}

	err = rh.Redis.HSet("job_reg_"+jobId, "t_step", len(model.Process)).Err()
	if err != nil {
		return
	}

	for i, v := range model.Process {
		value, err := json.Marshal(v)
		if err != nil {
			return err
		}

		err = rh.Redis.HSet("job_reg_"+jobId, fmt.Sprintf("step_%d", i), value).Err()
		if err != nil {
			return err
		}
	}

	// job注册信息 24h 过期
	err = rh.Redis.Expire("job_reg_"+jobId, time.Hour*24).Err()
	if err != nil {
		return
	}
	return
}
