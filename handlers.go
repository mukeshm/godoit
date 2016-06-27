package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func myIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Proto, r.RemoteAddr, r.UserAgent())
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
