package PhMessage

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhSchemaModel"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	kafkaAvro "github.com/elodina/go-kafka-avro"
)

type PhKafkaHandler struct {
	schemaRepositoryUrl	string
}

func (handler PhKafkaHandler) New(srUrl string) *PhKafkaHandler{
	return &PhKafkaHandler{schemaRepositoryUrl:srUrl}
}

func (handler PhKafkaHandler) Send(topic string, model PhSchemaModel.PhAvroModel) (err error) {
	record, err := model.GenSchema(model).GenRecord(model)
	if err != nil {
		return
	}

	encoder := kafkaAvro.NewKafkaAvroEncoder(handler.schemaRepositoryUrl)
	recordByteArr, err := encoder.Encode(record)
	if err != nil {
		return
	}

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		return
	}

	bkc.Produce(&topic, recordByteArr)
	return
}
