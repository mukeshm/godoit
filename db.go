package main

import "errors"

var tasks Tasks = make(Tasks)

func AddTask(t Task) error {
	tasks[t.Id] = t
	return nil
}

func GetTask(id string) (Task, error) {
	val, ok := tasks[id]
	if ok {
		return val, nil
	}
	return val, errors.New("Error getting task")
}

func DeleteTask(id string) error {
	delete(tasks, id)
	return nil
}

func GetTasks() Tasks {
	return tasks
}

func UpdateTask(id string, t Task) error {
	val, ok := tasks[id]
	if ok {
		val.Completed = t.Completed
		val.Due = t.Due
		val.Id = id
		val.Title = t.Title
		tasks[id] = val
		return nil
	}
	return errors.New("failed to update task : task does not exist")
}
