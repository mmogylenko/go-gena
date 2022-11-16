package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/phuslu/log"
)

const (
	addr = ":8080"
)

var (
	logger = &log.Logger{
		TimeFormat: "15:04:05",
		Caller:     2,
		Writer: &log.ConsoleWriter{
			ColorOutput: true,
			QuoteString: true,
		},
		Level: log.InfoLevel,
	}
)

func main() {

	router := mux.NewRouter()
	router.Use(LOGMiddleware)

	router.HandleFunc("/health", StatusHandler)

	api := router.PathPrefix("/api").Subrouter()
	api.Use(JSONMiddleware)
	api.HandleFunc("/info", InfoHandler)
	api.HandleFunc("/calculate", CalculateHandler)
	api.HandleFunc("/calculate/{number}", CalculateHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	service := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.
		Info().
		Str("addr", addr).
		Msg("Listening on")

	logger.
		Fatal().
		Err(service.ListenAndServe()).
		Msg("Failed to start service")

}
