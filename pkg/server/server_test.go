package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req,err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)

	handler.ServeHTTP(rr,req)

	// Test status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: status code expected %v got %v", http.StatusOK, rr.Code)
	}
}

func TestStatus(t *testing.T) {
	// Prepare route to test
	req,err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test status code
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Status)

	handler.ServeHTTP(rr,req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: status code got %v want %v", status, http.StatusOK)
	}

	// Test response body
	expected := `{"status":"ok"}`
	if rr.Body.String() != expected {
		t.Errorf("Error: expected %v got %v", expected, rr.Body.String())
	}
	
}