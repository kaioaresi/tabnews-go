package main

import (
	"log"
	"net/http"
	"tabnews-go/pkg/db"
	"tabnews-go/pkg/logger"
	"tabnews-go/pkg/server"
)

func main() {
	mux := http.NewServeMux()

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal("Failed to create logger", err)
	}

	db, err := db.NewDBClient(logger)
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()

	serverConfig, err := server.NewServerConfig(db, logger)
	if err != nil {
		logger.Error(err)
	}

	mux.HandleFunc("/", serverConfig.Home)
	mux.HandleFunc("/api/v1/status", serverConfig.Status)
	mux.HandleFunc("/api/v1/migrations", serverConfig.Migrations)

	logger.Info("Server listing :8080....")
	logger.Error(http.ListenAndServe(":8080", mux))

}
