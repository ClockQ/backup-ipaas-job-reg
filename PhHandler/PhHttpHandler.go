package PhHandler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMessage"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"io/ioutil"
	"net/http"
)

func PhHttpHandler(kh *PhMessage.PhKafkaHandler, _ *PhMessage.PhMqttHandler) (_ func(http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response []byte

		switch r.Method {
		case "POST":
			var model PhModel.JobReg
			body, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(body, &model)
			if err != nil {
				response = []byte("Query Error: " + err.Error())
			}

			jobRequest := PhModel.JobRequest{}.GenTestData()
			// TODO topic 写死了
			err = kh.Send("cjob-test", jobRequest)
			if err != nil {
				response = []byte("Kafka Send Error: " + err.Error())
			}

			response = []byte("The call is successful")
		default:
			response = []byte("Bad Request Method")
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		w.Write(response)
	}
}
