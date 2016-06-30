package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func myIndexHandler(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "indexhandler")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func tasksIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	printLog(*r, "tasksIndex")
	ts := getTasks()
	fmt.Fprintf(w, "%v", ts)
}

//return individual task
func showTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "showTask")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	t, err := GetTask(vars["taskID"])
	if err != nil {
		fmt.Fprintf(w, "%+v", errorToJSON(err))
	} else {
		fmt.Fprintf(w, "%v", taskToJSON(t))
	}
}

//insert task into tasks
func insertTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "insertTask")
	w.Header().Set("Content-Type", "application/json")
	t, err := createTask(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%+v", errorToJSON(err))
	} else {
		AddTask(t)
		fmt.Fprintf(w, "%+v", taskToJSON(t))
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "deleteTask")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := DeleteTask(vars["taskID"])
	if err != nil {
		fmt.Fprintf(w, "%+v", errorToJSON(err))
	} else {
		fmt.Fprintf(w, "%v", successToJSON("Task "+vars["taskID"]+" removed"))
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "updateTask")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	t, e := jsonToTask(r.Body)
	if e != nil {
		fmt.Fprintf(w, "%+v", errorToJSON(e))
	} else {
		err := UpdateTask(vars["taskID"], t)
		if err != nil {
			fmt.Fprintf(w, "%+v", errorToJSON(err))
		} else {
			fmt.Fprintf(w, "%v", successToJSON("Task "+vars["taskID"]+" updated"))
		}
	}
}

func printLog(r http.Request, name string) {
	log.Println(r.Method, r.URL, r.Proto, r.RemoteAddr, name)
}
