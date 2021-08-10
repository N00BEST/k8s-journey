package main

import (
	"encoding/json"
	"net/http"
)

func readinessProbe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(true)
}

func livenessProbe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(true)
}