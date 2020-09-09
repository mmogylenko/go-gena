package main

import (
	"encoding/json"
	"fmt"
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
	result := calculate()
	writeJSON(writer, fmt.Sprintf("x = %v", result))
}
