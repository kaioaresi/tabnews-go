package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"tabnews-go/pkg/db"
)

type ServerConfig struct{}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}

func (s ServerConfig) Home(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Home page!")
}

func (s ServerConfig) Status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	_, err := db.Ping()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "mensage": "error on check db!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status":"ok", "db_version":"1.1.1"}`)

}

func (s ServerConfig) Migrations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `[]`)
}
