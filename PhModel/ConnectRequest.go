package PhModel

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

type ConnectRequest struct {
	*PhSchemaModel
	JobId        string
	Tag          string // TM | UCB | MAX
	SourceConfig string
	SinkConfig   string
}

func (model ConnectRequest) New() *ConnectRequest {
	model.PhSchemaModel = &PhSchemaModel{}
	return &model
}

func (model *ConnectRequest) Inject(data map[string]interface{}) error {
	sinkConfif, _ := json.Marshal(data["SinkConfig"])
	data["SinkConfig"] = string(sinkConfif)
	sourceConfif, _ := json.Marshal(data["SourceConfig"])
	data["SourceConfig"] = string(sourceConfif)

	err := mapstructure.Decode(data, model)
	return err
}
