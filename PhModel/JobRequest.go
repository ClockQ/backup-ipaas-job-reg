package PhModel

type JobRequest struct {
	*PhSchemaModel
	JobType        string // Jar | Py | R
	Class          string
	Master         string
	DeployMode     string
	ExecutorMemory string
	ExecutorCores  string
	NumExecutors   string
	Queue          string
	Target         string
	File           string
	Parameter      string
}

// TODO: 注意删除
func (model JobRequest) GenTestData() *JobRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.JobType = "Jar"
	model.Class = "com.pharbers.ipaas.data.driver.Main"
	model.Master = "yarn"
	model.DeployMode = "cluster"
	model.ExecutorMemory = "1G"
	model.ExecutorCores = "1"
	model.NumExecutors = "2"
	model.Queue = "default"
	model.Target = "hdfs:///jars/context/job-context.jar"
	model.File = "hdfs:///jars/context/pharbers_config/kafka_config.xml,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.truststore.jks,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.keystore.jks"
	model.Parameter = "yaml hdfs hdfs:///test/MZclean.yaml"
	return &model
}
