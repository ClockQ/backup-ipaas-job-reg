package PhJobManager

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/hashicorp/go-uuid"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

const (
	MqttUrl             = "http://59.110.31.215:6542/v0/publish"
	MqttChannel         = "test-qi/"
	SchemaRepositoryUrl = "http://123.56.179.133:8081"
	RedisHost           = "59.110.31.215"
	RedisPort           = "6378"
	RedisPwd            = ""
)

func TestPhJobReg(t *testing.T) {
	kh := PhHelper.PhKafkaHelper{}.New(SchemaRepositoryUrl)
	mh := PhHelper.PhMqttHelper{}.New(MqttUrl, MqttChannel)
	rh := PhHelper.PhRedisHelper{}.New(RedisHost, RedisPort, RedisPwd)

	Convey("测试 TM JobReg", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := []interface{}{
			map[string]interface{}{"a": "001"},
			map[string]interface{}{"a": "002"},
			map[string]interface{}{"a": "003"},
		}
		model := PhModel.JobReg{JobId: jobId, Process: process}

		err := PhJobReg(model, kh, mh, rh)
		So(err, ShouldBeNil)

		cStepStr, _ := rh.Redis.HGet(jobId, "c_step").Result()
		cStep, _ := strconv.Atoi(cStepStr)
		So(cStep, ShouldEqual, 0)

		tStepStr, _ := rh.Redis.HGet(jobId, "t_step").Result()
		tStep, _ := strconv.Atoi(tStepStr)
		So(tStep, ShouldEqual, 3)

		for i := 0; i < tStep; i++ {
			stepStr, _ := rh.Redis.HGet(jobId, "step_"+string(i)).Result()
			So(stepStr, ShouldNotBeNil)
		}
	})
}
