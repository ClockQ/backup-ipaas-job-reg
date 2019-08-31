package PhModel

import (
	"github.com/mitchellh/mapstructure"
)

type TmAggRequest struct {
	*PhSchemaModel
	JobId      string
	RequestId  string // UUID
	ProposalId string
	ProjectId  string
	PeriodId   string
	Phase      int32
	Strategy   string // Agg2Cal | Cal2Report | Report2Show
}

func (model TmAggRequest) New() *TmAggRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	return &model
}

func (model *TmAggRequest) Inject(data map[string]interface{}) error {
	err := mapstructure.Decode(data, model)
	return err
}
