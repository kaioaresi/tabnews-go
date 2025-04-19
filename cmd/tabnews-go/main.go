package main

import (
	"log"
	"net/http"
	"tabnews-go/pkg/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.Home)
	mux.HandleFunc("/api/v1/status", server.Status)
	mux.HandleFunc("/api/v1/migrations", server.Migrations)

	log.Println("Server listing :8080....")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
