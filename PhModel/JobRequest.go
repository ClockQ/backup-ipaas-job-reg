package PhModel

type JobRequest struct {
	*PhSchemaModel
	Name           string
	JobType        string // Jar | Py | R
	Class          string
	Master         string
	DeployMode     string
	ExecutorMemory string
	ExecutorCores  string
	NumExecutors   string
	Queue          string
	Target         string
	Files          string
	Parameters     string
}

// TODO: 注意删除
func (model JobRequest) GenTestData() *JobRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.Name = "TestJob"
	model.JobType = "Jar"
	model.Class = "com.pharbers.ipaas.data.driver.Main"
	model.Master = "yarn"
	model.DeployMode = "cluster"
	model.ExecutorMemory = "1G"
	model.ExecutorCores = "1"
	model.NumExecutors = "2"
	model.Queue = "default"
	model.Target = "hdfs:///jars/context/job-context.jar"
	model.Files = "hdfs:///jars/context/pharbers_config/kafka_config.xml,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.truststore.jks,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.keystore.jks"
	model.Parameters = "yaml hdfs hdfs:///test/MZclean.yaml"
	return &model
}

// TODO: 注意删除
func (model JobRequest) GenTMData() *JobRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	model.Name = "TM Calc"
	model.JobType = "R"
	model.Master = "yarn"
	model.DeployMode = "cluster"
	model.ExecutorMemory = "1G"
	model.ExecutorCores = "1"
	model.NumExecutors = "2"
	model.Queue = "researches"
	model.Target = "hdfs:///jars/context/NTMR/TMUCBCalMain.R"
	model.Files = "hdfs:///jars/context/pharbers_config/kafka_config.xml,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.truststore.jks,hdfs:///jars/context/pharbers_config/secrets/kafka.broker1.keystore.jks"
	model.Parameters = "hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/cal_data hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/weightages hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/curves-n hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/competitor"
	return &model
}
