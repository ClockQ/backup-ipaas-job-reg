package PhModel

type JobProcess struct {
	PsType  string
	Actions map[string]interface{}
}

type JobReg struct {
	JobId   string
	Process []JobProcess
}
