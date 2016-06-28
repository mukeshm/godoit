package main

import (
	"encoding/json"
	"errors"
	"github.com/mukeshm/uuid"
	"io"
	"time"
)

func createTask(r io.Reader) (Task, error) {
	var task Task
	err := json.NewDecoder(r).Decode(&task)
	if err != nil {
		return task, err
	}
	if task.Title == "" {
		return task, errors.New("Title cannot be empty")
	}
	task.Kind = "task"
	u, _ := uuid.GenerateV4()
	uuid := u.String()
	task.Id = uuid
	var dt time.Time
	if task.Due == dt {
		task.Due = time.Now().AddDate(0, 0, 1)
	}
	return task, nil
}

func taskToJSON(t Task) string {
	bs, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(bs)
}

func getTasks() string {
	ba, err := json.Marshal(tasks)
	if err != nil {
		return ""
	}
	return string(ba)
}