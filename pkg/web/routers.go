package web

import (
	"log"
	"net/http"
	"tabnews-go/config"
	"tabnews-go/pkg/db"
	"tabnews-go/pkg/logger"
)

func Routers(cfg *config.DBCredentials, dbConfig *db.DBConfig) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal("Failed to create logger", err)
		return nil, err
	}

	serverConfig, err := NewServerConfig(dbConfig, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	mux.HandleFunc("/", serverConfig.Home)
	mux.HandleFunc("/api/v1/status", serverConfig.Status)
	mux.HandleFunc("/api/v1/migrations", serverConfig.Migrations)

	return mux, nil
}
