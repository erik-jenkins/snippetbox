package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erik-jenkins/snippetbox/pkg/middleware"
)

func main() {
	port := 4000
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.Get(home))
	mux.HandleFunc("/snippet", middleware.Get(showSnippet))
	mux.HandleFunc("/snippet/create", middleware.Post(createSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
