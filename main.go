package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a new HTTP server to listen for requests
	http.Handle("/metrics", promhttp.Handler())

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

// Handler function to extract and encode metrics as JSON
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the metrics from the Prometheus server
	metrics, err := promhttp.DefaultGatherer.Gather()
	if err != nil {
		http.Error(w, "Error gathering metrics", http.StatusInternalServerError)
		return
	}

	// Encode the metrics as JSON and send them in the response
	encoder := json.NewEncoder(w)
	err = encoder.Encode(metrics)
	if err != nil {
		http.Error(w, "Error encoding metrics as JSON", http.StatusInternalServerError)
		return
	}
}
