package main

import "net/http"

type Route struct {
	Name string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"/api",
		myIndexHandler,
	},
}