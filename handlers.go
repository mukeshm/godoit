package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func myIndexHandler(w http.ResponseWriter, r *http.Request) {
	printLog(r, "indexhandler")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func tasksIndex(w http.ResponseWriter, r *http.Request) {
	printLog(r, "tasksIndex")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func showTask(w http.ResponseWriter, r *http.Request) {
	printLog(r, "showTask")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func insertTask(w http.ResponseWriter, r *http.Request) {
	printLog(r, "insertTask")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	printLog(r, "deleteTask")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	printLog(r, "updateTask")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func printLog(r http.Request, name string){
	log.Println(r.Method, r.URL, r.Proto, r.RemoteAddr, r.UserAgent(), name)
}