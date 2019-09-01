package PhModel

type JobRegResponse struct {
	JobId  string
	Status string // RUNNING, FINISH, ERROR
	Error  interface{}
	Result interface{}
}

func (model JobRegResponse) SetRunning(jobId string, result interface{}) *JobRegResponse {
	model.JobId = jobId
	model.Status = "RUNNING"
	model.Result = result
	return &model
}

func (model JobRegResponse) SetFinish(jobId string, result interface{}) *JobRegResponse {
	model.JobId = jobId
	model.Status = "FINISH"
	model.Result = result
	return &model
}

func (model JobRegResponse) SetError(jobId string, errMsg interface{}) *JobRegResponse {
	model.JobId = jobId
	model.Status = "ERROR"
	model.Error = errMsg
	return &model
}
