package server

import (
	"encoding/json"
	"log"
	"net/http"
	"tabnews-go/pkg/db"
)

type ServerConfig struct {
	dbInfos *db.DbInfo
}

func NewServerConfig(dbInfos db.DBAccess) (*ServerConfig, error) {

	infosDB, err := dbInfos.GetDBInfos()
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		dbInfos: infosDB,
	}, nil
}

func (s ServerConfig) Home(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home tabnews"))
}

func (s ServerConfig) Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(s.dbInfos)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "error encode"})
		return
	}
}

func (s ServerConfig) Migrations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte("Migrations page"))
}
