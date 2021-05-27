package TaskModel

import "encoding/json"

type TaskModel struct {
	RegisterTime  int64
	ExecutionTime int64
	Args          []byte
	TaskType      int
}

func NewTaskModel(registerTime int64, executionTime int64, args map[string]interface{}, taskType int) *TaskModel {
	var data []byte
	data, err := json.Marshal(args)
	if err != nil {
		return nil
	}
	return &TaskModel{
		RegisterTime:  registerTime,
		ExecutionTime: executionTime,
		Args:          data,
		TaskType:      taskType,
	}
}
