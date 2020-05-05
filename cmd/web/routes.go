package main

import (
	"net/http"

	"github.com/erik-jenkins/snippetbox/pkg/middleware"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.Get(app.home))
	mux.HandleFunc("/snippet", middleware.Get(app.showSnippet))
	mux.HandleFunc("/snippet/create", middleware.Post(app.createSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
