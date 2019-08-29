package main

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHandler"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMessage"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"os"
)

const (
	LogPath             = "job_reg.log"
	KafkabRokerUrl      = "123.56.179.133:9092"
	SchemaRepositoryUrl = "http://123.56.179.133:8081"
	KafkaGroup          = "test20190828"
	CaLocation          = "/opt/kafka/pharbers-secrets/snakeoil-ca-1.crt"
	CASignedLocation    = "/opt/kafka/pharbers-secrets/kafkacat-ca1-signed.pem"
	SSLKeyLocation      = "/opt/kafka/pharbers-secrets/kafkacat.client.key"
	SSLPwd              = "pharbers"
	MqttUrl				= "http://59.110.31.215:6542/v0/publish"
)

func main() {
	_ = os.Setenv("LOG_PATH", LogPath)
	_ = os.Setenv("BM_KAFKA_BROKER", KafkabRokerUrl)
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", SchemaRepositoryUrl)
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", KafkaGroup)
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", CaLocation)
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", CASignedLocation)
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", SSLKeyLocation)
	_ = os.Setenv("BM_KAFKA_SSL_PASS", SSLPwd)

	jobRequest := PhModel.JobRequest{}.GenTestData()
	err := PhMessage.PhKafkaHandler{}.New(SchemaRepositoryUrl).Send("cjob-test", jobRequest)
	bmerror.PanicError(err)

	jobResponse := PhModel.JobResponse{}
	err = PhMessage.PhKafkaHandler{}.New(SchemaRepositoryUrl).
		Linster([]string{"cjob-test2"}, &jobResponse, PhHandler.JobResponseHandler)
	bmerror.PanicError(err)

	model := jobResponse //PhModel.JobState{}.GenTestData()
	err = PhMessage.PhMqttHandler{}.New(MqttUrl).Send("test-qi/", model)
	bmerror.PanicError(err)
}

func send() {
	err := PhMessage.PhKafkaHandler{}.New(SchemaRepositoryUrl).Send("cjob-test", PhModel.JobRequest{}.GenTestData())
	bmerror.PanicError(err)
}

func linster() {
	err := PhMessage.PhKafkaHandler{}.New(SchemaRepositoryUrl).
		Linster([]string{"cjob-test2"}, &(PhModel.JobResponse{}), PhHandler.JobResponseHandler)
	bmerror.PanicError(err)
}

func sendMqtt() {
	model := PhModel.JobState{}.GenTestData()
	err := PhMessage.PhMqttHandler{}.New(MqttUrl).Send("test-qi/", model)
	bmerror.PanicError(err)
}
