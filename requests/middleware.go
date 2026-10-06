package main

import (
	"log/slog"
	"net/http"
)

// Middleware functions for HTTP requests and responses.

// SetHeader is a middleware function that sets the Content-Type and X-Server headers for all HTTP responses.
func SetHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-Server", "Go Server")
		next.ServeHTTP(w, r)
	})
}

// LogRequest is a middleware function that logs the details of incoming HTTP requests, including the client's IP address, request method, URL, and protocol version.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)

		userAttrs := slog.Group("user", slog.String("ip", ip))
		requestAttrs := slog.Group("request", slog.String("method", method), slog.String("url", url), slog.String("proto", proto))

		slog.Info("Request Received:", userAttrs, requestAttrs)
		next.ServeHTTP(w, r)
	})
}
