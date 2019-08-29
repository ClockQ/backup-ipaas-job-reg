package PhMessage

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"log"
	"net/http"
)

type PhMqttHandler struct {
	Url string
}

func (handler PhMqttHandler) New(url string) *PhMqttHandler {
	return &PhMqttHandler{Url: url}
}

func (handler *PhMqttHandler) Send(channel string, model PhModel.PhMessageModel) (err error) {
	log.Printf("MQTT 发送消息 %s 到 %s \n", model, channel)

	header := make(map[string]string)
	header["method"] = "Publish"
	header["channel"] = channel
	header["topic"] = ""

	body := make(map[string]interface{})
	body["header"] = header
	body["payload"] = model

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", handler.Url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("MQTT 发送失败")
	}
	return
}
