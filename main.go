package main

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhSchemaModel"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"os"
)

const (
	KafkabRokerUrl      = "123.56.179.133:9092"
	SchemaRepositoryUrl = "http://123.56.179.133:8081"
	KafkaGroup          = "test20190828"
	CaLocation          = "/opt/kafka/pharbers-secrets/snakeoil-ca-1.crt"
	CASignedLocation    = "/opt/kafka/pharbers-secrets/kafkacat-ca1-signed.pem"
	SSLKeyLocation      = "/opt/kafka/pharbers-secrets/kafkacat.client.key"
	SSLPwd              = "pharbers"
)

func main() {
	_ = os.Setenv("BM_KAFKA_BROKER", KafkabRokerUrl)
	_ = os.Setenv("BM_KAFKA_SCHEMA_REGISTRY_URL", SchemaRepositoryUrl)
	_ = os.Setenv("BM_KAFKA_CONSUMER_GROUP", KafkaGroup)
	_ = os.Setenv("BM_KAFKA_CA_LOCATION", CaLocation)
	_ = os.Setenv("BM_KAFKA_CA_SIGNED_LOCATION", CASignedLocation)
	_ = os.Setenv("BM_KAFKA_SSL_KEY_LOCATION", SSLKeyLocation)
	_ = os.Setenv("BM_KAFKA_SSL_PASS", SSLPwd)

	model := PhSchemaModel.JobResponse{}.GenTestData()
	record, err := model.GenRecord(&model)
	bmerror.PanicError(err)

	encoder := kafkaAvro.NewKafkaAvroEncoder(SchemaRepositoryUrl)
	recordByteArr, err := encoder.Encode(record)
	bmerror.PanicError(err)

	bkc, err := bmkafka.GetConfigInstance()
	bmerror.PanicError(err)
	topic := "cjob-test"
	bkc.Produce(&topic, recordByteArr)
}
