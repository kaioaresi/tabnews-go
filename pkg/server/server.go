package server

import (
	"fmt"
	"io"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,"Home page!")
}

func Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	w.WriteHeader(http.StatusOK)
	io.WriteString(w,`{"status":"ok"}`)
}

func Migrations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	
	w.WriteHeader(http.StatusOK)
	io.WriteString(w,`[]`)
}
