package PhModel

import (
	"github.com/mitchellh/mapstructure"
)

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
	Jars           string
	Files          string
	Conf           string
	Target         string
	Parameters     string
}

func (model JobRequest) New() *JobRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	return &model
}

func (model *JobRequest) Inject(data map[string]interface{}) error {
	err := mapstructure.Decode(data, model)
	return err
}
