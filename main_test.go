package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler(t *testing.T) {
	// Reset start time for consistent testing
	startTime = time.Now()

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware(healthHandler))

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json",
			contentType)
	}

	// Parse and validate JSON response
	var response HealthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	// Validate response structure
	if response.Status != "ok" {
		t.Errorf("handler returned wrong status: got %v want ok", response.Status)
	}

	if response.Uptime == "" {
		t.Error("handler returned empty uptime")
	}

	if response.Version == "" {
		t.Error("handler returned empty version")
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware(helloHandler))

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json",
			contentType)
	}

	// Parse and validate JSON response
	var response MessageResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	// Validate response payload
	expectedMessage := "Hello, World! Welcome to the Go web app."
	if response.Message != expectedMessage {
		t.Errorf("handler returned wrong message: got %v want %v",
			response.Message, expectedMessage)
	}
}

func TestMetricsHandler(t *testing.T) {
	// Reset counters
	requestCount = 0
	healthCount = 0
	helloCount = 0
	metricsCount = 0

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware(metricsHandler))

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check content type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want application/json",
			contentType)
	}

	// Parse and validate JSON response
	var response MetricsResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	// Validate response structure
	if response.Version == "" {
		t.Error("handler returned empty version")
	}

	if response.Uptime == "" {
		t.Error("handler returned empty uptime")
	}

	// Check that metrics endpoint was called (should be at least 1)
	if response.Requests.Metrics < 1 {
		t.Error("metrics counter should be at least 1")
	}
}
