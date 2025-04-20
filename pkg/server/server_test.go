package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	serverConfigClient := NewServerConfig()
	// Prepare route to test
	req, err := http.NewRequest("GET", "/api/v1/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test status code
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serverConfigClient.Status)

	handler.ServeHTTP(rr, req)

	// Test response body
	expected := `{"status":"ok", "db_version":"1.1.1"}`
	if rr.Body.String() != expected {
		t.Errorf("Error: expected %v got %v", expected, rr.Body.String())
	}

}

func TestMigrations(t *testing.T) {
	serverConfigClient := NewServerConfig()
	req, err := http.NewRequest("GET", "/api/v1/migrations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serverConfigClient.Migrations)
	handler.ServeHTTP(rr, req)

	expected := `[]`
	if rr.Body.String() != expected {
		t.Errorf("Error: expected %v got %v", expected, rr.Body.String())
	}
}

func TestHandlersStatusCode(t *testing.T) {

	serverConfigClient := NewServerConfig()

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
