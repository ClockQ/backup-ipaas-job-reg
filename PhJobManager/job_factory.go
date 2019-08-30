package PhJobManager

//
//var PhJobFuncFactory = map[string]func(*PhHelper.PhKafkaHelper, *PhModel.JobProcess) error{
//	"JOB": func(kh *PhHelper.PhKafkaHelper, model *PhModel.JobProcess) error {
//		jobRequestTopic := os.Getenv("JOB_REQUEST_TOPIC")
//		err = kh.Send(jobRequestTopic, model)
//		return nil
//	},
//}
//
//func GetJobFunc(jobType string) func(*PhHelper.PhKafkaHelper, *PhModel.JobProcess) error {
//	return PhJobFuncFactory[strings.ToUpper(jobType)]
//}
