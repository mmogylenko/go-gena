package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type InfoResponse struct {
	Title    string `json:"title"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
}

// CalculateHandler is One-shot math operation (square root)
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var (
		payload map[string]interface{}
	)

	vars := mux.Vars(r)
	if vars["number"] != "" {
		if s, err := strconv.ParseFloat(vars["number"], 32); err == nil {
			payload = map[string]interface{}{
				"message": fmt.Sprintf("The square root of %s is %f", vars["number"], SquareRootOf(s)),
			}
		}
	} else {
		payload = map[string]interface{}{
			"message": SquareRoot(),
		}
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error().Err(err).Msg("Failed to encode JSON")
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"status": "OK"}); err != nil {
		logger.Error().Err(err).Msg("Failed to encode JSON")
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		logger.Error().Err(err).Msg("Host name reported by the kernel")
	}

	response := InfoResponse{
		Title:    "Home | go-gena",
		Hostname: host,
		Message:  fmt.Sprintf("Hello from %s", AppName),
		Version:  AppVersion,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error().Err(err).Msg("Failed to encode JSON")
	}
}
