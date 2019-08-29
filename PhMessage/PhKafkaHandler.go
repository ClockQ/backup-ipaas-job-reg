package PhMessage

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"log"
)

type PhKafkaHandler struct {
	schemaRepositoryUrl string
}

func (handler PhKafkaHandler) New(srUrl string) *PhKafkaHandler {
	return &PhKafkaHandler{schemaRepositoryUrl: srUrl}
}

func (handler PhKafkaHandler) Send(topic string, model PhModel.PhAvroModel) (err error) {
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

func (handler PhKafkaHandler) Linster(topics []string, msgModel interface{}, subscribeFunc func(receive interface{})) error {

	bkc, err := bmkafka.GetConfigInstance()
	if err != nil {
		return err
	}

	bkc.SubscribeTopics(topics, func(receive interface{}) {
		decoder := kafkaAvro.NewKafkaAvroDecoder(handler.schemaRepositoryUrl)
		record, err := decoder.Decode(receive.([]byte))
		if err != nil {
			errMsg := fmt.Sprintf("接受的 %s 信息解析出错: %s", topics, err)
			log.Println(errMsg)
			bmlog.StandardLogger().Error(errMsg)
			return
		}

		err = json.Unmarshal([]byte(record.(*avro.GenericRecord).String()), msgModel)
		if err != nil {
			errMsg := fmt.Sprintf("接受的 %s 信息解析出错: %s", topics, err)
			log.Println(errMsg)
			bmlog.StandardLogger().Error(errMsg)
			return
		}

		subscribeFunc(msgModel)
	})

	return nil
}
