package PhJobManager

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhChannel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhEnv"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/hashicorp/go-uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestProcessExec_Channel_M2H(t *testing.T) {
	PhEnv.SetEnv()

	kh := PhChannel.PhKafkaHelper{}.New(SchemaRepositoryUrl)

	Convey("测试 TM Channel: Mongodb -> HDFS", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := PhModel.JobProcess{
			PsType: "CHANNEL",
			Actions: map[string]interface{}{
				"jobId": jobId,
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
					"connector.class":    "io.confluent.connect.hdfs.HdfsSinkConnector",
					"tasks.max":          "1",
					"jobId":              jobId,
					"topics":             jobId,
					"hdfs.url":           "hdfs://192.168.100.137:9000/logs/testqi/parquet/",
					"format.class":       "io.confluent.connect.hdfs.parquet.ParquetFormat",
					"rotate.interval.ms": "1000",
					"flush.size":         "10",
				},
			},
		}

		err := processExec(&process, kh)
		So(err, ShouldBeNil)
	})
}

func TestProcessExec_Channel_H2M(t *testing.T) {
	PhEnv.SetEnv()

	kh := PhChannel.PhKafkaHelper{}.New(SchemaRepositoryUrl)

	Convey("测试 TM Channel: HDFS -> Mongodb", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := PhModel.JobProcess{
			PsType: "CHANNEL",
			Actions: map[string]interface{}{
				"jobId": jobId,
				"Tag":   "TM",
				"SourceConfig": map[string]interface{}{
					"connector.class": "com.github.mmolimar.kafka.connect.fs.FsSourceConnector",
					"tasks.max":       "1",

					"jobId":   "abc0000010",
					"topic":   "abc0000010",
					"fs.uris": "hdfs://192.168.100.137:9000/test/UCBTest/inputParquet/TMInputParquet0820/output/264e12ff-62a5-4cdf-bec5-2eb2014f6154/cal_report/",

					"file_reader.class": "com.github.mmolimar.kafka.connect.fs.file.reader.ParquetFileReader",
					"policy.class":      "com.github.mmolimar.kafka.connect.fs.policy.SimplePolicy",
					"policy.recursive":  "true",
					"policy.regexp":     ".*",
				},
				"SinkConfig": map[string]interface{}{
					"connector.class":                     "at.grahsl.kafka.connect.mongodb.MongoDbSinkConnector",
					"tasks.max":                           "1",
					"job":                                 jobId,
					"topic":                               jobId,
					"mongodb.connection.uri":              "mongodb://192.168.100.176:27017/job_reg_test?w=1&journal=true",
					"mongodb.collection":                  "job_reg_cal_report",
					"key.converter":                       "io.confluent.connect.avro.AvroConverter",
					"key.converter.schema.registry.url":   "http://59.110.31.50:8081",
					"value.converter":                     "io.confluent.connect.avro.AvroConverter",
					"value.converter.schema.registry.url": "http://59.110.31.50:8081",
					"connection.timeout.ms":               "5000",
				},
			},
		}

		err := processExec(&process, kh)
		So(err, ShouldBeNil)
	})
}

func TestProcessExec_Channel_M2E(t *testing.T) {
	PhEnv.SetEnv()

	kh := PhChannel.PhKafkaHelper{}.New(SchemaRepositoryUrl)

	Convey("测试 TM Channel: Mongodb -> ES", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := PhModel.JobProcess{
			PsType: "CHANNEL",
			Actions: map[string]interface{}{
				"jobId": jobId,
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
					"connection.url":        "http://192.168.100.176:9200",
					"type.name":             "",
					"read.timeout.ms":       "10000",
					"connection.timeout.ms": "5000",
				},
			},
		}

		err := processExec(&process, kh)
		So(err, ShouldBeNil)
	})
}

func TestProcessExec_Job(t *testing.T) {
	PhEnv.SetEnv()

	kh := PhChannel.PhKafkaHelper{}.New(SchemaRepositoryUrl)

	Convey("测试 TM JobExec", t, func() {
		jobId, _ := uuid.GenerateUUID()
		process := PhModel.JobProcess{
			PsType: "JOB",
			Actions: map[string]interface{}{
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
		err := processExec(&process, kh)
		So(err, ShouldBeNil)
	})
}
