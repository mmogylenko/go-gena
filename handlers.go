package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, msg string) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": msg}); err != nil {
		w.WriteHeader(503)
		log.Fatal(err)
	}

}

func start(writer http.ResponseWriter, request *http.Request) {
	//start load generator gouroutine
	go load(ctx)
	writeJSON(writer, "Started Load Generator")
}

func stop(writer http.ResponseWriter, request *http.Request) {
	// cancel ctx
	cancel()
	writeJSON(writer, "Stopping Load Generator")
}
