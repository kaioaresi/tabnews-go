package main

import (
	"log"
	"net/http"
	"tabnews-go/pkg/server"
)

func main() {
	http.HandleFunc("/", server.Home)

	log.Println("Starting web server...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
