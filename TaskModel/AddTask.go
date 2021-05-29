package TaskModel

import (
	"github.com/northcity0406/delayTask/DBHandler"
	"encoding/json"
	"time"
)

func AddTask(execTime int64, args map[string]interface{}, taskType int) error {
	task := NewTaskModel(time.Now().Unix(), execTime, args, taskType)
	data, _ := json.Marshal(task)
	return DBHandler.ZAddValue(task, data, float64(execTime))
}
