package server

import (
	"net/http"
	"net/http/httptest"
	"tabnews-go/internal/db"
	"testing"
)

func TestHandlersStatusCode(t *testing.T) {
	db, err := db.NewDBClient()
	if err != nil {
		t.Errorf("Error db connection %v", err)
	}
	defer db.Close()

	serverConfigClient, err := NewServerConfig(db)
	if err != nil {
		t.Errorf("Error server %v", err)
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
				t.Fatalf("Error to create request %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.HandlerF)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.Status {
				t.Errorf("Test %s - error: expected %v got %v", tt.Name, tt.Status, rr.Code)
			}
		})
	}
}
