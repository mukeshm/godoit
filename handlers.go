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
	printLog(*r, "tasksIndex")
	ts := getTasks()
	fmt.Fprintf(w, "%v", ts)
}

//return individual task
func showTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "showTask")
	vars := mux.Vars(r)
	t, err := GetTask(vars["taskID"])
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
	} else {
		fmt.Fprintf(w, "%v", taskToJSON(t))
	}
}

//insert task into tasks
func insertTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "insertTask")
	t, err := createTask(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
	} else {
		AddTask(t)
		fmt.Printf("%v\n", t)
		fmt.Fprintf(w, "%+v", t)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "deleteTask")
	vars := mux.Vars(r)
	err := DeleteTask(vars["taskID"])
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
	} else {
		fmt.Fprintf(w, "%v removed", vars["taskID"])
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	printLog(*r, "updateTask")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func printLog(r http.Request, name string) {
	log.Println(r.Method, r.URL, r.Proto, r.RemoteAddr, r.UserAgent(), name)
}
