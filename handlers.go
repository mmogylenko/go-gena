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

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.New("index.html").ParseFiles(path.Join("web", "index.html"))
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(path.Join("web", "index.html") + err.Error()))
// 		return
// 	}
// 	data := struct {
// 		Title string
// 		Logo  string
// 	}{
// 		Title: ,
// 		Logo:  "https://user-images.githubusercontent.com/7536624/92551931-42128880-f214-11ea-8ebb-a71817168353.png",
// 	}

// 	if err := tmpl.Execute(w, data); err != nil {
// 		http.Error(w, path.Join("web", "index.html")+err.Error(), http.StatusInternalServerError)
// 	}

// }
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
