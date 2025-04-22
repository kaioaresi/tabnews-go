package server

import (
	"encoding/json"
	"log"
	"net/http"
	"tabnews-go/internal/logger"
	"tabnews-go/pkg/db"
	"time"
)

const (
	contentTypeHeader = "Content-type"
	applicationJson   = "application/json"
)

type ServerConfig struct {
	UpdateAt     time.Time          `json:"update_at"`
	Dependencies DependenciesStatus `json:"dependencies"`
	logger       *logger.Logger
}

type DependenciesStatus struct {
	Database *db.DbInfo `json:"database"`
}

func NewServerConfig(dbInfos db.DBAccess) (*ServerConfig, error) {

	lg, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	infosDB, err := dbInfos.GetDBInfos()
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		UpdateAt: time.Now(),
		logger:   lg,
		Dependencies: DependenciesStatus{
			Database: infosDB,
		},
	}, nil
}

func (s ServerConfig) Home(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home tabnews"))
}

func (s ServerConfig) Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(contentTypeHeader, applicationJson)
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		s.logger.Errorf("Error encoding error message in Status endpoint", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "error encode"})
		return
	}
}

func (s ServerConfig) Migrations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(contentTypeHeader, applicationJson)
	w.Write([]byte("Migrations page"))
}
