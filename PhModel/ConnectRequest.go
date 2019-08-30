package PhModel

import "encoding/json"

type ConnectRequest struct {
	*PhSchemaModel
	JobId        string
	Tag          string // TM | UCB | MAX
	SourceConfig string
	SinkConfig   string
}

type SourceConfig map[string]string

type SinkConfig map[string]string

// TODO: 注意删除
func (model ConnectRequest) GenTestData() *ConnectRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.JobId = "abc0000010"
	model.Tag = "TM"
	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTMSourceMongoData("abc0000010"))
	model.SourceConfig = string(sourceConfigBytes)
	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTestData())
	model.SinkConfig = string(sinkConfigBytes)
	return &model
}

// TODO: 注意删除
func (model SourceConfig) GenTMSourceMongoData(jobId string) SourceConfig {
	model["connector.class"] = "com.pharbers.kafka.connect.mongodb.MongodbSourceConnector"
	model["tasks.max"] = "1"
	model["job"] = jobId
	model["topic"] = jobId
	model["connection"] = "mongodb://192.168.100.176:27017"
	model["database"] = "test"
	model["collection"] = "PhAuth"
	model["filter"] = "{}"
	return model
}

// TODO: 注意删除
func (model SinkConfig) GenTestData() SinkConfig {
	model["connector.class"] = "com.pharbers.kafka.connect.elasticsearch.ElasticsearchSinkConnector"
	model["tasks.max"] = "1"
	model["jobId"] = "abc0000010"
	model["topics"] = "abc0000010"
	model["key.ignore"] = "true"
	model["connection.url"] = "http://59.110.31.215:9200"
	model["type.name"] = ""
	model["read.timeout.ms"] = "10000"
	model["connection.timeout.ms"] = "5000"
	return model
}

// TODO: 注意删除
func (model ConnectRequest) GenTMHDFS2MongoData() *ConnectRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.JobId = "abc0000010"
	model.Tag = "TM"
	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTMSourceHdfsData())
	model.SourceConfig = string(sourceConfigBytes)
	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTM2MongoData())
	model.SinkConfig = string(sinkConfigBytes)
	return &model
}

// TODO: 注意删除
func (model SourceConfig) GenTMSourceHdfsData() SourceConfig {
	model["connector.class"] = "com.github.mmolimar.kafka.connect.fs.FsSourceConnector"
	model["tasks.max"] = "1"

	model["jobId"] = "abc0000010"
	model["topic"] = "abc0000010"
	model["fs.uris"] = "hdfs://192.168.100.137:9000/test/UCBTest/inputParquet/TMInputParquet0820/output/264e12ff-62a5-4cdf-bec5-2eb2014f6154/cal_report/"

	model["file_reader.class"] = "com.github.mmolimar.kafka.connect.fs.file.reader.ParquetFileReader"
	model["policy.class"] = "com.github.mmolimar.kafka.connect.fs.policy.SimplePolicy"
	model["policy.recursive"] = "true"
	model["policy.regexp"] = ".*"

	return model
}

// TODO: 注意删除
func (model SinkConfig) GenTM2MongoData() SinkConfig {
	model["connector.class"] = "at.grahsl.kafka.connect.mongodb.MongoDbSinkConnector"
	model["tasks.max"] = "1"

	model["jobId"] = "abc0000010"
	model["topics"] = "abc0000010"
	model["mongodb.connection.uri"] = "mongodb://192.168.100.176:27017/job_reg_test?w=1&journal=true"
	model["mongodb.collection"] = "job_reg_cal_report"

	model["key.converter"] = "io.confluent.connect.avro.AvroConverter"
	model["key.converter.schema.registry.url"] = "http://59.110.31.50:8081"
	model["value.converter"] = "io.confluent.connect.avro.AvroConverter"
	model["value.converter.schema.registry.url"] = "http://59.110.31.50:8081"
	model["connection.timeout.ms"] = "5000"

	return model
}

// TODO: 注意删除
func (model ConnectRequest) GenTMMongo2HDFSData(jobId string) *ConnectRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.JobId = jobId
	model.Tag = "TM"
	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTMSourceMongoData(jobId))
	model.SourceConfig = string(sourceConfigBytes)
	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTM2HdfsData(jobId))
	model.SinkConfig = string(sinkConfigBytes)
	return &model
}

// TODO: 注意删除
func (model SinkConfig) GenTM2HdfsData(jobId string) SinkConfig {
	model["connector.class"] = "io.confluent.connect.hdfs.HdfsSinkConnector"
	model["tasks.max"] = "1"

	model["jobId"] = jobId
	model["topics"] = jobId
	model["hdfs.url"] = "hdfs://192.168.100.137:9000/logs/testqi/parquet/"

	model["format.class"] = "io.confluent.connect.hdfs.parquet.ParquetFormat"
	model["rotate.interval.ms"] = "1000"
	model["flush.size"] = "10"

	return model
}
