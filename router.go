package main

import "github.com/gorilla/mux"

func newRouter() *mux.Router {
	router := mux.NewRouter()
	for _, v := range routes {
		router.Methods(v.Method).Path(v.Pattern).HandlerFunc(v.HandlerFunc)
	}
	return router
}
