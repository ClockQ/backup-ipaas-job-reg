package PhJobManager

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/hashicorp/go-uuid"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
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

	_ = os.Setenv("IS_TEST", "true")
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

//// TODO: 注意删除
//func (model JobRequest) GenTestData() *JobRequest {
//	model.PhSchemaModel = &PhSchemaModel{}
//	model.Name = "TestJob"
//	model.JobType = "Jar"
//	model.Class = "com.pharbers.ipaas.data.driver.Main"
//	model.Master = "yarn"
//	model.DeployMode = "cluster"
//	model.ExecutorMemory = "1G"
//	model.ExecutorCores = "1"
//	model.NumExecutors = "2"
//	model.Queue = "default"
//	model.Target = "hdfs:///jars/context/job-context.jar"
//	model.Files = "hdfs:///jars/context/pharbers_config/kafka_config.xml,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.truststore.jks,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.keystore.jks"
//	model.Parameters = "yaml hdfs hdfs:///test/MZclean.yaml"
//	return &model
//}
//
//// TODO: 注意删除
//func (model JobRequest) GenTMData() *JobRequest {
//	model.PhSchemaModel = &PhSchemaModel{}
//	model.Name = "TmCalc"
//	model.JobType = "R"
//	model.Master = "yarn"
//	model.DeployMode = "cluster"
//	model.ExecutorMemory = "1G"
//	model.ExecutorCores = "1"
//	model.NumExecutors = "2"
//	model.Queue = "researches"
//	model.Target = "hdfs:///jars/context/NTMR/TMUCBCalMain.R"
//	model.Files = "hdfs:///jars/context/NTMR/AddCols.R,hdfs:///jars/context/NTMR/CastCol2Double.R,hdfs:///jars/context/NTMR/ColMin.R,hdfs:///jars/context/NTMR/ColMax.R,hdfs:///jars/context/NTMR/ColRename.R,hdfs:///jars/context/NTMR/ColSum.R,hdfs:///jars/context/NTMR/CurveFunc.R,hdfs:///jars/context/NTMR/UCBDataBinding.R,hdfs:///jars/context/NTMR/TMCalCurveSkeleton2.R,hdfs:///jars/context/NTMR/UCBCalFuncs.R,hdfs:///jars/context/NTMR/TMCalResAchv.R"
//	model.Parameters = "hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/cal_data hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/weightages hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/curves-n hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/competitor"
//	return &model
//}

//
//// TODO: 注意删除
//func (model ConnectRequest) GenTestData(jobId string) *ConnectRequest {
//	model.PhSchemaModel = &PhSchemaModel{}
//	model.JobId = jobId
//	model.Tag = "TM"
//	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTMSourceMongoData(jobId))
//	model.SourceConfig = string(sourceConfigBytes)
//	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTestData(jobId))
//	model.SinkConfig = string(sinkConfigBytes)
//	return &model
//}
//
//// TODO: 注意删除
//func (model SourceConfig) GenTMSourceMongoData(jobId string) SourceConfig {
//	model["connector.class"] = "com.pharbers.kafka.connect.mongodb.MongodbSourceConnector"
//	model["tasks.max"] = "1"
//	model["job"] = jobId
//	model["topic"] = jobId
//	model["connection"] = "mongodb://192.168.100.176:27017"
//	model["database"] = "test"
//	model["collection"] = "PhAuth"
//	model["filter"] = "{}"
//	return model
//}
//
//// TODO: 注意删除
//func (model SinkConfig) GenTestData(jobId string) SinkConfig {
//	model["connector.class"] = "com.pharbers.kafka.connect.elasticsearch.ElasticsearchSinkConnector"
//	model["tasks.max"] = "1"
//	model["jobId"] = jobId
//	model["topics"] = jobId
//	model["key.ignore"] = "true"
//	model["connection.url"] = "http://59.110.31.215:9200"
//	model["type.name"] = ""
//	model["read.timeout.ms"] = "10000"
//	model["connection.timeout.ms"] = "5000"
//	return model
//}
//
//// TODO: 注意删除
//func (model ConnectRequest) GenTMHDFS2MongoData() *ConnectRequest {
//	model.PhSchemaModel = &PhSchemaModel{}
//	model.JobId = "abc0000010"
//	model.Tag = "TM"
//	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTMSourceHdfsData())
//	model.SourceConfig = string(sourceConfigBytes)
//	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTM2MongoData())
//	model.SinkConfig = string(sinkConfigBytes)
//	return &model
//}
//
//// TODO: 注意删除
//func (model SourceConfig) GenTMSourceHdfsData() SourceConfig {
//	model["connector.class"] = "com.github.mmolimar.kafka.connect.fs.FsSourceConnector"
//	model["tasks.max"] = "1"
//
//	model["jobId"] = "abc0000010"
//	model["topic"] = "abc0000010"
//	model["fs.uris"] = "hdfs://192.168.100.137:9000/test/UCBTest/inputParquet/TMInputParquet0820/output/264e12ff-62a5-4cdf-bec5-2eb2014f6154/cal_report/"
//
//	model["file_reader.class"] = "com.github.mmolimar.kafka.connect.fs.file.reader.ParquetFileReader"
//	model["policy.class"] = "com.github.mmolimar.kafka.connect.fs.policy.SimplePolicy"
//	model["policy.recursive"] = "true"
//	model["policy.regexp"] = ".*"
//
//	return model
//}
//
//// TODO: 注意删除
//func (model SinkConfig) GenTM2MongoData() SinkConfig {
//	model["connector.class"] = "at.grahsl.kafka.connect.mongodb.MongoDbSinkConnector"
//	model["tasks.max"] = "1"
//
//	model["jobId"] = "abc0000010"
//	model["topics"] = "abc0000010"
//	model["mongodb.connection.uri"] = "mongodb://192.168.100.176:27017/job_reg_test?w=1&journal=true"
//	model["mongodb.collection"] = "job_reg_cal_report"
//
//	model["key.converter"] = "io.confluent.connect.avro.AvroConverter"
//	model["key.converter.schema.registry.url"] = "http://59.110.31.50:8081"
//	model["value.converter"] = "io.confluent.connect.avro.AvroConverter"
//	model["value.converter.schema.registry.url"] = "http://59.110.31.50:8081"
//	model["connection.timeout.ms"] = "5000"
//
//	return model
//}
//
//// TODO: 注意删除
//func (model ConnectRequest) GenTMMongo2HDFSData(jobId string) *ConnectRequest {
//	model.PhSchemaModel = &PhSchemaModel{}
//	model.JobId = jobId
//	model.Tag = "TM"
//	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTMSourceMongoData(jobId))
//	model.SourceConfig = string(sourceConfigBytes)
//	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTM2HdfsData(jobId))
//	model.SinkConfig = string(sinkConfigBytes)
//	return &model
//}
//
//// TODO: 注意删除
//func (model SinkConfig) GenTM2HdfsData(jobId string) SinkConfig {
//	model["connector.class"] = "io.confluent.connect.hdfs.HdfsSinkConnector"
//	model["tasks.max"] = "1"
//
//	model["jobId"] = jobId
//	model["topics"] = jobId
//	model["hdfs.url"] = "hdfs://192.168.100.137:9000/logs/testqi/parquet/"
//
//	model["format.class"] = "io.confluent.connect.hdfs.parquet.ParquetFormat"
//	model["rotate.interval.ms"] = "1000"
//	model["flush.size"] = "10"
//
//	return model
//}

func TestProcessExec_Channel(t *testing.T) {
	setEnv()

	kh := PhHelper.PhKafkaHelper{}.New(SchemaRepositoryUrl)

	Convey("测试 TM Channel", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := PhModel.JobProcess{
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
			},
		}

		err := ProcessExec(&process, kh)
		So(err, ShouldBeNil)
	})
}

func TestProcessExec_Job(t *testing.T) {
	setEnv()

	kh := PhHelper.PhKafkaHelper{}.New(SchemaRepositoryUrl)

	Convey("测试 TM JobExec", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := PhModel.JobProcess{
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
			},
		}
		err := ProcessExec(&process, kh)
		So(err, ShouldBeNil)
	})
}
