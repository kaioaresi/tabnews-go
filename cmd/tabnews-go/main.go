package main

import (
	"log"
	"net/http"
	"tabnews-go/config"
	"tabnews-go/pkg/db"
	"tabnews-go/pkg/logger"
	"tabnews-go/pkg/web"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal("Failed to create logger", err)
	}
	defer logger.Sugar.Sync()

	db, err := db.NewDBClient(logger, config.StringConnection())
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()

	mux, err := web.Routers(config, db)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Server listing :8080....")
	logger.Error(http.ListenAndServe(":8080", mux))
}
