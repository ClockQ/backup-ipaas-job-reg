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
	Url     string
	Channel string
}

func (handler PhMqttHandler) New(url, channel string) *PhMqttHandler {
	return &PhMqttHandler{Url: url, Channel: channel}
}

func (handler *PhMqttHandler) SetChannel(url string) *PhMqttHandler {
	handler.Url = url
	return handler
}

func (handler *PhMqttHandler) Send(model PhModel.PhMessageModel) (err error) {
	log.Printf("MQTT 发送消息 %#v 到 %s \n", model, handler.Channel)

	header := make(map[string]string)
	header["method"] = "Publish"
	header["channel"] = handler.Channel
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
