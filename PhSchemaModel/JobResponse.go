package PhSchemaModel

type JobResponse struct {
	PhSchemaModel
	JobType string // Jar | Py | R
	Class string
	Master string
	DeployMode string
	ExecutorMemory string
	ExecutorCores string
	NumExecutors string
	Queue string
	Target string
	Parameter string
}

// TODO: 注意删除
func (model JobResponse) GenTestData() JobResponse {
	model.JobType = "Jar"
	model.Class = "com.pharbers.ipaas.data.driver.Main"
	model.Master = "yarn"
	model.DeployMode = "cluster"
	model.ExecutorMemory = "1G"
	model.ExecutorCores = "1"
	model.NumExecutors = "2"
	model.Queue = "default"
	model.Target = "hdfs:///jars/context/job-context.jar"
	model.Parameter = "yaml hdfs hdfs:///test/MZclean.yaml"
	return model
}
