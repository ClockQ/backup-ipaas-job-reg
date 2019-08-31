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

const (
	JobRequestTopic      = "cjob-test"
	JobResponseTopic     = "cjob-test2"
	ConnectRequestTopic  = "ConnectRequest"
	ConnectResponseTopic = "ConnectResponse"
)

func setEnv() {
	const (
		Ip           = ""
		Port         = "9213"
		Prefix       = "/v1.0/job_reg"
		LogPath      = "job_reg.log"
		WriteTimeout = "4"

		JobRequestTopic      = "cjob-test"
		JobResponseTopic     = "cjob-test2"
		ConnectRequestTopic  = "ConnectRequest"
		ConnectResponseTopic = "ConnectResponse"

		MqttUrl     = "http://59.110.31.215:6542/v0/publish"
		MqttChannel = "test-qi/"

		RedisHost = "59.110.31.215"
		RedisPort = "6378"
		RedisPwd  = ""

		KafkabRokerUrl      = "123.56.179.133:9092"
		SchemaRepositoryUrl = "http://123.56.179.133:8081"
		KafkaGroup          = "test20190828"
		CaLocation          = "/opt/kafka/pharbers-secrets/snakeoil-ca-1.crt"
		CASignedLocation    = "/opt/kafka/pharbers-secrets/kafkacat-ca1-signed.pem"
		SSLKeyLocation      = "/opt/kafka/pharbers-secrets/kafkacat.client.key"
		SSLPwd              = "pharbers"
	)

	_ = os.Setenv("JOB_REQUEST_TOPIC", JobRequestTopic)
	_ = os.Setenv("JOB_RESPONSE_TOPIC", JobResponseTopic)
	_ = os.Setenv("CONNECT_REQUEST_TOPIC", ConnectRequestTopic)
	_ = os.Setenv("CONNECT_RESPONSE_TOPIC", ConnectResponseTopic)

	_ = os.Setenv("MQTT_URL", MqttUrl)
	_ = os.Setenv("MQTT_CHANNEL", MqttChannel)

	_ = os.Setenv("REDIS_HOST", RedisHost)
	_ = os.Setenv("REDIS_PORT", RedisPort)
	_ = os.Setenv("REDIS_PWD", RedisPwd)

	_ = os.Setenv("BM_KAFKA_BROKER", KafkabRokerUrl)
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", SchemaRepositoryUrl)
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", KafkaGroup)
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", CaLocation)
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", CASignedLocation)
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", SSLKeyLocation)
	_ = os.Setenv("BM_KAFKA_SSL_PASS", SSLPwd)
}

func TestPhJobExec_Channel(t *testing.T) {
	_ = os.Setenv("IS_TEST", "true")
	_ = os.Setenv("JOB_REQUEST_TOPIC", JobRequestTopic)
	_ = os.Setenv("CONNECT_REQUEST_TOPIC", ConnectRequestTopic)

	kh := PhHelper.PhKafkaHelper{}.New(SchemaRepositoryUrl)
	mh := PhMqttHelper.PhMqttHelper{}.New(MqttUrl, MqttChannel)
	rh := PhHelper.PhRedisHelper{}.New(RedisHost, RedisPort, RedisPwd)

	Convey("测试 TM Channel", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := []PhModel.JobProcess{{
			PsType: "CHANNEL",
			JobConfig: map[string]interface{}{
				"JobId": jobId,
				"Tag":   "TM",
				"SourceConfig": map[string]interface{}{
					"connector.class": "com.pharbers.kafka.connect.mongodb.MongodbSourceConnector",
					"tasks.max":       "1",
					"job":             jobId,
					"topic":           jobId,
					"connection":      "mongodb://192.168.100.176:27017",
					"database":        "test",
					"collection":      "PhAuth",
					"filter":          "{}",
				},
				"SinkConfig": map[string]interface{}{
					"connector.class":       "com.pharbers.kafka.connect.elasticsearch.ElasticsearchSinkConnector",
					"tasks.max":             "1",
					"jobId":                 jobId,
					"topics":                jobId,
					"key.ignore":            "true",
					"connection.url":        "http://59.110.31.215:9200",
					"type.name":             "",
					"read.timeout.ms":       "10000",
					"connection.timeout.ms": "5000",
				},
			}},
		}
		model := PhModel.JobReg{JobId: jobId, Process: process}

		err := PhJobReg(model, kh, mh, rh)
		So(err, ShouldBeNil)

		beforeExecStepStr, _ := rh.Redis.HGet(jobId, "c_step").Result()
		beforeExecStep, _ := strconv.Atoi(beforeExecStepStr)
		So(beforeExecStep, ShouldEqual, 0)

		PhJobExec(jobId, kh, mh, rh)

		rh.Redis.Del(jobId)
	})
}

func TestPhJobExec_Job(t *testing.T) {
	setEnv()

	kh := PhHelper.PhKafkaHelper{}.New(SchemaRepositoryUrl)
	mh := PhMqttHelper.PhMqttHelper{}.New(MqttUrl, MqttChannel)
	rh := PhHelper.PhRedisHelper{}.New(RedisHost, RedisPort, RedisPwd)

	Convey("测试 TM JobExec", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := []PhModel.JobProcess{{
			PsType: "JOB",
			JobConfig: map[string]interface{}{
				"Name":           "TM-Submit",
				"JobType":        "R",
				"Master":         "yarn",
				"DeployMode":     "cluster",
				"ExecutorMemory": "1G",
				"ExecutorCores":  "1",
				"NumExecutors":   "1",
				"Queue":          "default",
				"Files":          "/Users/qianpeng/GitHub/NTMRCal/NTM/TMCalColFuncs.R,/Users/qianpeng/GitHub/NTMRCal/NTM/TMCalProcess.R,/Users/qianpeng/GitHub/NTMRCal/NTM/TMCalResAchv.R,/Users/qianpeng/GitHub/NTMRCal/NTM/TMDataCbind.R",
				"Conf":           "spark.yarn.appMasterEnv.KAFKA_PROXY_URI=http://59.110.31.50:8082/topics,spark.yarn.appMasterEnv.KAFKA_PROXY_R_CAL_TOPIC=listeningJobTask",
				"Target":         "/Users/qianpeng/GitHub/NTMRCal/main.R",
				"Parameters":     "NTM hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/cal_data hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/weightages hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/manager hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/curves-n hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/competitor hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/standard_time hdfs://192.168.100.137:9000//test/TMTest/inputParquet/TMInputParquet0815/level_data " + jobId,
			}},
		}
		model := PhModel.JobReg{JobId: jobId, Process: process}

		err := PhJobReg(model, kh, mh, rh)
		So(err, ShouldBeNil)

		beforeExecStepStr, _ := rh.Redis.HGet(jobId, "c_step").Result()
		beforeExecStep, _ := strconv.Atoi(beforeExecStepStr)
		So(beforeExecStep, ShouldEqual, 0)

		PhJobExec(jobId, kh, mh, rh)

		rh.Redis.Del(jobId)
	})
}
