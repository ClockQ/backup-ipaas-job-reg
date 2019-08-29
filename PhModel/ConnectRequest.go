package PhModel

import "encoding/json"

type ConnectRequest struct {
	*PhSchemaModel
	JobId        string
	Tag          string // TM | UCB | MAX
	SourceConfig string
	SinkConfig   string
}

// TODO: 注意删除
func (model ConnectRequest) GenTestData() *ConnectRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.JobId = "abc0000004"
	model.Tag = "TM"
	sourceConfigBytes, _ := json.Marshal(SourceConfig{}.GenTestData())
	model.SourceConfig = string(sourceConfigBytes)
	sinkConfigBytes, _ := json.Marshal(SinkConfig{}.GenTestData())
	model.SinkConfig = string(sinkConfigBytes)
	return &model
}

type SourceConfig map[string]string

// TODO: 注意删除
func (model SourceConfig) GenTestData() SourceConfig {
	model["connector.class"] = "com.pharbers.kafka.connect.mongodb.MongodbSourceConnector"
	model["tasks.max"] = "1"
	model["job"] = "abc0000004"
	model["topic"] = "abc0000004"
	model["connection"] = "mongodb://192.168.100.176:27017"
	model["database"] = "test"
	model["collection"] = "PhAuth"
	model["filter"] = "{}"
	return model
}

type SinkConfig map[string]string

// TODO: 注意删除
func (model SinkConfig) GenTestData() SinkConfig {
	model["connector.class"] = "com.pharbers.kafka.connect.elasticsearch.ElasticsearchSinkConnector"
	model["tasks.max"] = "1"
	model["jobId"] = "abc0000004"
	model["topics"] = "abc0000004"
	model["key.ignore"] = "true"
	model["connection.url"] = "http://59.110.31.215:9200"
	model["type.name"] = ""
	model["read.timeout.ms"] = "10000"
	model["connection.timeout.ms"] = "5000"
	return model
}
