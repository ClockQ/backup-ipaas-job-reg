package PhHandler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhHelper"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhJobManager"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhMqttHelper"
	"io/ioutil"
	"net/http"
)

func PhHttpHandler(kh *PhHelper.PhKafkaHelper,
	mh *PhMqttHelper.PhMqttHelper,
	rh *PhHelper.PhRedisHelper) (_ func(http.ResponseWriter, *http.Request)) {

	return func(w http.ResponseWriter, r *http.Request) {
		var response []byte

		switch r.Method {
		case "POST":
			var model PhModel.JobReg
			body, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(body, &model)
			if err != nil {
				response = []byte("Param Parse Error: " + err.Error())
				break
			}

			// 进行 Job 注册
			err = PhJobManager.PhJobReg(model, kh, mh, rh)
			if err != nil {
				response = []byte("Job Reg Error: " + err.Error())
				break
			}

			// 协程开始执行 Job
			go PhJobManager.JobExec(model.JobId, kh, mh, rh)

			response = []byte("The Call Is Successful")
		default:
			response = []byte("Bad Request Method")
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("Content-Type", "application/json")             //返回数据格式是json
		w.Write(response)
	}
}
