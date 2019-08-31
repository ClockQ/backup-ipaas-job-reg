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
	Files          string
	Conf           string
	Target         string
	Parameters     string
}

func (model JobRequest) New() *JobRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	return &model
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
	model.Name = "TmCalc"
	model.JobType = "R"
	model.Master = "yarn"
	model.DeployMode = "cluster"
	model.ExecutorMemory = "1G"
	model.ExecutorCores = "1"
	model.NumExecutors = "2"
	model.Queue = "researches"
	model.Target = "hdfs:///jars/context/NTMR/TMUCBCalMain.R"
	model.Files = "hdfs:///jars/context/NTMR/AddCols.R,hdfs:///jars/context/NTMR/CastCol2Double.R,hdfs:///jars/context/NTMR/ColMin.R,hdfs:///jars/context/NTMR/ColMax.R,hdfs:///jars/context/NTMR/ColRename.R,hdfs:///jars/context/NTMR/ColSum.R,hdfs:///jars/context/NTMR/CurveFunc.R,hdfs:///jars/context/NTMR/UCBDataBinding.R,hdfs:///jars/context/NTMR/TMCalCurveSkeleton2.R,hdfs:///jars/context/NTMR/UCBCalFuncs.R,hdfs:///jars/context/NTMR/TMCalResAchv.R"
	model.Parameters = "hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/cal_data hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/weightages hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/curves-n hdfs://192.168.100.137:9000//test/UCBTest/inputParquet/TMInputParquet0820/competitor"
	return &model
}
