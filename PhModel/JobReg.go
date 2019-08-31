package PhModel

type JobReg struct {
	Id      string
	JobId   string
	Process []JobProcess
}

type JobProcess struct {
	PsType    string
	JobConfig map[string]interface{}
}
