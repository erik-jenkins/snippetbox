package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	standardMiddleware := []Middleware{
		app.RecoverPanic,
		app.LogRequests,
		SecureHeaders,
	}

	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/snippet/{id:[0-9]+}", app.showSnippet).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippetForm).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippet).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	return ChainMiddleware(r, standardMiddleware...)
}
