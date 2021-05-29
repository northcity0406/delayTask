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
	fmt.Println("pong")
	return nil
}

func ECHO(data []byte) error {
	fmt.Println("echo")
	return nil
}

func LoopTasks() {
	for {
		taskList, _ := DBHandler.ZRangeValue(task, 0, 4)
		for _, val := range taskList {
			taskModel := TaskModel.TaskModel{}
			err := json.Unmarshal([]byte(val), &taskModel)
			if !taskModel.CanExecute() {
				continue
			}
			if err == nil {
				err := tasks.DealTask(taskModel.TaskType, taskModel.Args)
				if err == nil {
					var rem []interface{}
					rem = append(rem, &taskModel)
					err = DBHandler.ZRemValueByMembers(task, rem)
					fmt.Println(err)
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
		execTime := time.Now().Add(5 * time.Second).Unix()
		if rand.Int()%3 == 0 {
			TaskModel.AddTask(execTime, args, tasks.Ping)
		} else {
			TaskModel.AddTask(execTime, args, tasks.ECHO)
		}
	}
	for {
	}
}
