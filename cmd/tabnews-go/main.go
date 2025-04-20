package main

import (
	"log"
	"net/http"
	"tabnews-go/pkg/server"
)

func main() {
	mux := http.NewServeMux()

	serverConfig := server.NewServerConfig()

	mux.HandleFunc("/", serverConfig.Home)
	mux.HandleFunc("/api/v1/status", serverConfig.Status)
	mux.HandleFunc("/api/v1/migrations", serverConfig.Migrations)

	log.Println("Server listing :8080....")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
