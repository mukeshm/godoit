package main

import "time"

type Task struct {
	Kind string
	Id string
	Title string
	Due time.Time
	completed bool
}

type Tasks map[string]Task