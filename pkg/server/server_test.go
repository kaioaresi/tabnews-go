package server

import (
	"net/http"
	"net/http/httptest"
	"tabnews-go/pkg/db"
	"testing"
)

func TestHandlersStatusCode(t *testing.T) {
	db, err := db.NewDBClient()
	if err != nil {
		t.Errorf("Error to db connection")
	}

	serverConfigClient, err := NewServerConfig(db)
	if err != nil {
		t.Errorf("Error server")
	}

	tests := []struct {
		Name     string
		Path     string
		Method   string
		Status   int
		HandlerF func(w http.ResponseWriter, req *http.Request)
	}{
		{
			Name:     "Home",
			Path:     "/",
			Method:   "GET",
			Status:   http.StatusOK,
			HandlerF: serverConfigClient.Home,
		},
		{
			Name:     "Status",
			Path:     "/api/v1/status",
			Method:   "GET",
			Status:   http.StatusOK,
			HandlerF: serverConfigClient.Status,
		},
		{
			Name:     "Migrations",
			Path:     "/api/v1/migrations",
			Method:   "GET",
			Status:   http.StatusOK,
			HandlerF: serverConfigClient.Migrations,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req, err := http.NewRequest(tt.Method, tt.Path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.HandlerF)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Error: expected %v got %v", http.StatusOK, rr.Code)
			}
		})
	}
}
