package main

import (
	"encoding/json"
	"fmt"
	"github.com/northcity0406/delayTask/DBHandler"
	"github.com/northcity0406/delayTask/TaskModel"
	"github.com/northcity0406/delayTask/tasks"
	"math/rand"
	"time"
)

const task = "tasks"

func Ping(data []byte) error {
	fmt.Println(string(data))
	fmt.Println("pong")
	return nil
}

func ECHO(data []byte) error {
	fmt.Println(string(data))
	return nil
}

func LoopTasks() {
	for {
		taskList, _ := DBHandler.ZRangeValue(task, 0, 4)
		for _, val := range taskList {
			taskModel := TaskModel.TaskModel{}
			err := json.Unmarshal([]byte(val), &taskModel)
			if err == nil {
				err := tasks.DealTask(taskModel.TaskType, taskModel.Args)
				if err == nil {
					var rem []interface{}
					rem = append(rem, val)
					_ = DBHandler.ZRemValueByMembers(task, rem)
				}
			}
		}
	}
}

func main() {
	_ = tasks.RegisterTasks(tasks.Ping, Ping)
	_ = tasks.RegisterTasks(tasks.ECHO, ECHO)
	go LoopTasks()

	for i := 0; i < 1000; i++ {
		args := make(map[string]interface{})
		now := time.Now().Unix()
		execTime := time.Now().Add(5 * time.Second).Unix()
		if rand.Int()%3 == 0 {
			taskModel := TaskModel.NewTaskModel(now, execTime, args, tasks.Ping)
			data, _ := json.Marshal(taskModel)
			_ = DBHandler.ZAddValue(task, data, float64(execTime))
		} else {
			taskModel := TaskModel.NewTaskModel(now, execTime, args, tasks.ECHO)
			data, _ := json.Marshal(taskModel)
			_ = DBHandler.ZAddValue(task, data, float64(execTime))
		}
	}
}
