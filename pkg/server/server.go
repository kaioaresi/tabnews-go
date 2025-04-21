package server

import (
	"encoding/json"
	"log"
	"net/http"
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
}

type DependenciesStatus struct {
	Database *db.DbInfo `json:"database"`
}

func NewServerConfig(dbInfos db.DBAccess) (*ServerConfig, error) {

	infosDB, err := dbInfos.GetDBInfos()
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		UpdateAt: time.Now(),
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
		log.Printf("Error encoding error message in Status endpoint %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "error encode"})
		return
	}
}

func (s ServerConfig) Migrations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(contentTypeHeader, applicationJson)
	w.Write([]byte("Migrations page"))
}
