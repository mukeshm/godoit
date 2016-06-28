package main

import "net/http"

type Route struct {
	Pattern string
	Method string
	HandlerFunc http.HandlerFunc
	Name string
}

type Routes []Route

var routes = Routes{
	Route{
		"/api",
		http.MethodGet,
		myIndexHandler,
		"Index",
	},
	Route{
		"/api/tasks",
		http.MethodGet,
		tasksIndex,
		"getTasks",
	},
	Route{
		"/api/tasks",
		http.MethodPost,
		insertTask,
		"insertTask",
	},
	Route{
		"/api/tasks/{taskID}",
		http.MethodGet,
		showTask,
		"getTaskByID",
	},
	Route{
		"/api/tasks/{taskID}",
		http.MethodDelete,
		deleteTask,
		"deleteTaskByID",
	},
	Route{
		"/api/tasks/{taskID}",
		http.MethodPut,
		updateTask,
		"updateTaskByID",
	},
}