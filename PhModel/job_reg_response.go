package PhModel

type JobRegResponse struct {
	JobId    string `json:"jobId"`
	Type     string `json:"type"`
	Status   string // RUNNING, FINISH, ERROR
	Progress string
	Error    map[string]interface{}
	Result   map[string]interface{}
}

func (model JobRegResponse) SetRunning(jobId, progress string, result map[string]interface{}) *JobRegResponse {
	model.JobId = jobId
	model.Type = "Calc"
	model.Status = "RUNNING"
	model.Progress = progress
	model.Result = result
	return &model
}

func (model JobRegResponse) SetFinish(jobId string, result map[string]interface{}) *JobRegResponse {
	model.JobId = jobId
	model.Type = "Calc"
	model.Status = "FINISH"
	model.Progress = "100"
	model.Result = result
	return &model
}

func (model JobRegResponse) SetError(jobId string, errMsg string) *JobRegResponse {
	model.JobId = jobId
	model.Type = "Calc"
	model.Status = "ERROR"
	model.Progress = "100"
	model.Error = map[string]interface{}{
		"message": errMsg,
	}
	return &model
}
