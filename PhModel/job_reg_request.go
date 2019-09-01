package PhModel

type JobRegRequest struct {
	JobId             string
	ProcessLocation   string // STRING, LOCAL, HDFS(暂无), OSS
	ProcessConfigType string // YAL(暂无), YAML(暂无), JSON
	ProcessConfig     interface{}
	Replace           map[string]interface{}
}
