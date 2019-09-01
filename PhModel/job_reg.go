package PhModel

type JobProcess struct {
	PsType  string
	Actions map[string]interface{}
}

type JobReg struct {
	Id      string
	JobId   string
	Process []JobProcess
}
