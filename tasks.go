package main

import "time"

type Task struct {
	Kind      string    `json:"kind"`
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Due       time.Time `json:"due,omitempty"`
	Completed bool      `json:"completed"`
}

type Tasks map[string]Task

type ErrorStruct struct {
	Kind        string    `json:"kind"`
	Message     string    `json:"message"`
}