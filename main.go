package main

import (
	"bytes"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHandler"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMessage"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	LogPath             = "dl.log"
	KafkabRokerUrl      = "123.56.179.133:9092"
	SchemaRepositoryUrl = "http://123.56.179.133:8081"
	KafkaGroup          = "test20190828"
	CaLocation          = "/opt/kafka/pharbers-secrets/snakeoil-ca-1.crt"
	CASignedLocation    = "/opt/kafka/pharbers-secrets/kafkacat-ca1-signed.pem"
	SSLKeyLocation      = "/opt/kafka/pharbers-secrets/kafkacat.client.key"
	SSLPwd              = "pharbers"
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

	//for i := 0; i < 10; i++ {
	//	send()
	//}
	//
	//linster()

	sendMqtt()
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
	url := "http://59.110.31.215:6542/v0/publish"

	var jsonStr = []byte(`{
	"header": {
		"method": "Publish",
		"channel": "test-qi/",
		"topic": ""
	},
	"payload": {
		"account": "demo",
		"progress": 10
	}
}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
