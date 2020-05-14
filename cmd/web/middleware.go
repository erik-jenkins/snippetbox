package main

import (
	"fmt"
	"net/http"
)

// Middleware defines a function that maps a handler to a handler
type Middleware func(http.Handler) http.Handler

// ChainMiddleware runs the middleware provided in `...Middleware` in reverse order
func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

// LogRequests logs the given request to the application's infolog
func (app *application) LogRequests(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// ContentTypeJSON sets the MIME type to json
func ContentTypeJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// SecureHeaders sets common security HTTP headers
func SecureHeaders(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// RecoverPanic makes sure a proper HTTP response is sent in the event of a panic
func (app *application) RecoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "Close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
