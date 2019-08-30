package PhJobManager

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"github.com/hashicorp/go-uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strconv"
	"testing"
)

func TestPhJobExec(t *testing.T) {
	_ = os.Setenv("IS_TEST", "true")

	kh := PhHelper.PhKafkaHelper{}.New(SchemaRepositoryUrl)
	mh := PhMqttHelper.PhMqttHelper{}.New(MqttUrl, MqttChannel)
	rh := PhHelper.PhRedisHelper{}.New(RedisHost, RedisPort, RedisPwd)

	Convey("测试 TM JobExec", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := []PhModel.JobProcess{
			{PsType: "Channel"},
		}
		model := PhModel.JobReg{JobId: jobId, Process: process}

		err := PhJobReg(model, kh, mh, rh)
		So(err, ShouldBeNil)

		beforeExecStepStr, _ := rh.Redis.HGet(jobId, "c_step").Result()
		beforeExecStep, _ := strconv.Atoi(beforeExecStepStr)
		So(beforeExecStep, ShouldEqual, 0)

		PhJobExec(jobId, kh, mh, rh)

		afterExecStepStr, _ := rh.Redis.HGet(jobId, "c_step").Result()
		afterExecStep, _ := strconv.Atoi(afterExecStepStr)
		So(afterExecStep, ShouldEqual, 1)

		rh.Redis.Del(jobId)
	})
}
