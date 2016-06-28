package main

import "errors"

var tasks Tasks

func AddTask(t Task) error {
	tasks[t.Id] = t
	return nil
}

func GetTask(id string) (Task, error){
	val, ok := tasks[id]
	if ok{
		return val, nil
	}
	return val, errors.New("Error getting task")
}

func DeleteTask(id string) error{
	delete(tasks, id)
	return  nil
}

func GetTasks() Tasks {
	return tasks
}