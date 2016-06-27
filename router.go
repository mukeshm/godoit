package main

import "net/http"

func newRouter() *http.ServeMux{
	router := http.NewServeMux()
	for _, v := range routes{
		router.HandleFunc(v.Name, v.HandlerFunc)
	}
	return router
}
