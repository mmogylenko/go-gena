package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/phuslu/log"
)

const (
	addr    = ":8080"
	timeout = 5
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
	api.HandleFunc("/load", CPULoadHandler)
	api.HandleFunc("/calculate", CalculateHandler)
	api.HandleFunc("/calculate/{number}", CalculateHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	service := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.
				Fatal().
				Err(err).
				Msg("Failed to start service")
		}
	}()

	logger.
		Info().
		Str("addr", addr).
		Msg("Listening on")

	<-done

	logger.
		Info().
		Int("timeout", timeout).
		Msg("Shutting down HTTP/HTTPS server")

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	if err := service.Shutdown(ctx); err != nil {
		logger.Fatal().Msgf("Server Shutdown Failed:%+v", err)
	}
}
