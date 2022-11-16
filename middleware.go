package main

import (
	"net/http"

	"golang.org/x/exp/maps"
)

func JSONMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func LOGMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.
			Info().
			Str("remote_addr", r.RemoteAddr).
			Strs("headers", maps.Keys(r.Header)).
			Str("path", r.URL.Path).
			Str("method", r.Method).
			Str("user-agent", r.UserAgent()).
			Msg("Incoming HTTP Request")

		next.ServeHTTP(w, r)
	})
}
