package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var (
	// Version is set at build time via -ldflags
	// Example: go build -ldflags "-X main.Version=v1.0.0"
	Version = "dev"

	// StartTime tracks when the server started for uptime calculation
	startTime = time.Now()

	// Request counters for metrics
	requestCount    int64
	healthCount     int64
	helloCount      int64
	metricsCount    int64
)

type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
	Version string `json:"version"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type MetricsResponse struct {
	Requests struct {
		Total  int64 `json:"total"`
		Health int64 `json:"health"`
		Hello  int64 `json:"hello"`
		Metrics int64 `json:"metrics"`
	} `json:"requests"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", middleware(healthHandler))

	// API endpoint
	mux.HandleFunc("/api/hello", middleware(helloHandler))

	// Metrics endpoint
	mux.HandleFunc("/metrics", middleware(metricsHandler))

	// Root endpoint
	mux.HandleFunc("/", rootHandler)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&healthCount, 1)

	uptime := time.Since(startTime)
	response := HealthResponse{
		Status:  "ok",
		Uptime:  fmt.Sprintf("%.0f", uptime.Seconds()),
		Version: Version,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&helloCount, 1)

	w.Header().Set("Content-Type", "application/json")
	response := MessageResponse{
		Message: "Hello, World! Welcome to the Go web app.",
	}
	json.NewEncoder(w).Encode(response)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&metricsCount, 1)

	uptime := time.Since(startTime)
	response := MetricsResponse{
		Requests: struct {
			Total   int64 `json:"total"`
			Health  int64 `json:"health"`
			Hello   int64 `json:"hello"`
			Metrics int64 `json:"metrics"`
		}{
			Total:   atomic.LoadInt64(&requestCount),
			Health:  atomic.LoadInt64(&healthCount),
			Hello:   atomic.LoadInt64(&helloCount),
			Metrics: atomic.LoadInt64(&metricsCount),
		},
		Uptime:  fmt.Sprintf("%.0f", uptime.Seconds()),
		Version: Version,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// middleware tracks request counts
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requestCount, 1)
		next(w, r)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Web App</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 30px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
        }
        .endpoint {
            background: #f8f9fa;
            padding: 10px;
            margin: 10px 0;
            border-left: 3px solid #007bff;
            font-family: monospace;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Go Web Application</h1>
        <p>Welcome to your Go web application!</p>
        <h2>Available Endpoints:</h2>
        <div class="endpoint">GET /health - Health check endpoint</div>
        <div class="endpoint">GET /api/hello - API hello endpoint</div>
        <div class="endpoint">GET /metrics - Metrics endpoint</div>
    </div>
</body>
</html>
`
	fmt.Fprintf(w, html)
}
