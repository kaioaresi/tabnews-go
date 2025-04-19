package main

import (
	"log"
	"net/http"
	"tabnews-go/pkg/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.Home)
	mux.HandleFunc("/status", server.Status)

	log.Println("Server listing :8080....")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
