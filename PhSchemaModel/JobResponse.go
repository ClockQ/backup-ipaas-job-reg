package PhSchemaModel

type JobResponse struct {
	*PhSchemaModel
	JobId        string
}

// TODO: 注意删除
func (model JobResponse) GenTestData() JobResponse {
	model.PhSchemaModel = &PhSchemaModel{}
	model.JobId = "test"
	return model
}
