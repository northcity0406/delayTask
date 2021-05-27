package tasks

import (
	"errors"
	"fmt"
	"sync"
)

type ConductTask func(data []byte) error

var Tasks map[int]ConductTask
var once sync.Once

func init() {
	once.Do(func() {
		Tasks = make(map[int]ConductTask)
	})
}

func RegisterTasks(key int, conduct ConductTask) error {
	_, ok := Tasks[key]
	if ok {
		return errors.New(fmt.Sprintf("The task type %d has been register!", key))
	}
	Tasks[key] = conduct
	return nil
}

func RemoveTasks(key int) error {
	_, ok := Tasks[key]
	if ok {
		delete(Tasks, key)
		return nil
	}
	return errors.New(fmt.Sprintf("The task type %d has been register!", key))
}

func DealTask(key int, data []byte) error {
	conduct, ok := Tasks[key]
	if !ok {
		return errors.New(fmt.Sprintf("The task type %d has been register!", key))
	}
	return conduct(data)
}
