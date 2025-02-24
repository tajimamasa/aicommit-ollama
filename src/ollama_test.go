package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckOllamaVersion(t *testing.T) {
	// Create a test server with a mock response
	mockResponse := `{"version": "0.5.1"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Call the function with the test server URL
	version, err := checkOllamaVersion(server.URL + "/")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the version
	expectedVersion := "0.5.1"
	if version != expectedVersion {
		t.Errorf("Expected version %s, got %s", expectedVersion, version)
	}
}

func TestCheckOllamaVersionError(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Call the function with the test server URL
	_, err := checkOllamaVersion(server.URL + "/")
	if err == nil {
		t.Fatal("Expected an error, got none")
	}
}
