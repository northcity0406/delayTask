package TaskModel

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack"
	"time"
)

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

func (t *TaskModel) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(t)
}

func (t *TaskModel) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, t)
}

func (t *TaskModel) CanExecute() bool {
	registerTime := time.Unix(t.ExecutionTime, 0)
	if registerTime.After(time.Now()) {
		return false
	}
	return true
}
