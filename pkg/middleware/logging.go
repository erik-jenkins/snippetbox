package middleware

import (
	"log"
	"net/http"
)

func LogRequests(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s", r.Method, r.URL.Path)
		handler(w, r)
	}
}
