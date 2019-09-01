package PhHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhChannel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhJobManager"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhModel"
	"github.com/PharbersDeveloper/ipaas-job-reg/PhThirdHelper"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func PhHttpHandler(kh *PhChannel.PhKafkaHelper,
	mh *PhThirdHelper.PhMqttHelper,
	rh *PhThirdHelper.PhRedisHelper) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var response []byte

		switch r.Method {
		case "POST":
			var model PhModel.JobRegRequest
			body, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(body, &model)
			if err != nil {
				resp := PhModel.JobRegResponse{}.SetError("", "Param Parse Error: "+err.Error())
				response, _ = json.Marshal(resp)
				break
			}

			// 提取 Job Process 信息
			jobReg := PhModel.JobReg{}
			jobReg.JobId = model.JobId
			jobReg.Process, err = extractProcess(model)
			if err != nil {
				resp := PhModel.JobRegResponse{}.SetError(model.JobId, "Extract Process Error: "+err.Error())
				response, _ = json.Marshal(resp)
				break
			}

			// 进行 Job 注册
			err = PhJobManager.PhJobReg(jobReg, kh, mh, rh)
			if err != nil {
				resp := PhModel.JobRegResponse{}.SetError(model.JobId, "Job Reg Error: "+err.Error())
				response, _ = json.Marshal(resp)
				break
			}

			// 协程开始执行 Job
			go PhJobManager.JobExec(model.JobId, kh, mh, rh)

			resp := PhModel.JobRegResponse{}.SetRunning(model.JobId, "0", map[string]interface{}{"message": "The Call Is Successful"})
			response, _ = json.Marshal(resp)
		default:
			resp := PhModel.JobRegResponse{}.SetError("", "Bad Request Method `"+r.Method+"`, Only Support `POST`")
			response, _ = json.Marshal(resp)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("Content-Type", "application/json")             //返回数据格式是json
		w.Write(response)
	}
}

func extractProcess(model PhModel.JobRegRequest) ([]PhModel.JobProcess, error) {
	// 从多种来源读入字节流
	var processBytes []byte
	switch strings.ToUpper(model.ProcessLocation) {
	case "STRING":
		pBytes, err := json.Marshal(model.ProcessConfig)
		if err != nil {
			return nil, err
		}
		processBytes = pBytes
	case "LOCAL":
		pBytes, err := ioutil.ReadFile(model.ProcessConfig.(string))
		if err != nil {
			return nil, err
		}
		processBytes = pBytes
	case "HDFS":
		return nil, errors.New("暂不支持 HDFS 上的文件")
	case "OSS":
		fallthrough
	default:
		endpoint := os.Getenv("ENDPOINT")
		accessKeyId := os.Getenv("ACCESS_KEY_ID")
		accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")

		oh := PhThirdHelper.PhOssHelper{}.New(endpoint, accessKeyId, accessKeySecret)
		path := strings.Split(model.ProcessConfig.(string), "/")
		str, err := oh.GetObject(path[0], path[1])
		if err != nil {
			return nil, err
		}
		processBytes = str
	}

	// 植入运行数据
	tmpStr := string(processBytes)
	for k, v := range model.Replace {
		tmpStr = strings.Replace(tmpStr, "##"+k+"##", fmt.Sprintf("%v", v), -1)
	}
	processBytes = []byte(tmpStr)

	// 解析字节流为 []PhModel.JobProcess
	processArray := make([]PhModel.JobProcess, 0)
	switch strings.ToUpper(model.ProcessConfigType) {
	case "YAL", "YAML":
		return nil, errors.New("暂不支持 YAML 格式的文件")
	case "JSON":
		fallthrough
	default:
		switch processBytes[0] {
		case 91: // '[' 的字节码
			err := json.Unmarshal(processBytes, &processArray)
			if err != nil {
				return nil, err
			}
		case 123: // '{' 的字节码
			process := PhModel.JobProcess{}
			err := json.Unmarshal(processBytes, &process)
			if err != nil {
				return nil, err
			}
			processArray = append(processArray, process)
		default:
			return nil, errors.New("不能处理的字节流")
		}
	}

	return processArray, nil
}
