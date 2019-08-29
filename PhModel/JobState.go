package PhModel

type JobState struct {
	Account  string
	Progress int
}

// TODO: 注意删除
func (model JobState) GenTestData() *JobState {
	model.Account = "demo"
	model.Progress = 10
	return &model
}
