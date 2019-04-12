package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/renderer/handlers"
)

func main() {
	// Router (Gorilla Mux)
	r := mux.NewRouter()

	// Takes care of Static file serving
	var dir string
	flag.StringVar(&dir, "dir", "./static", "usage")
	flag.Parse()

	// Static file server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// Dummy handler
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/blogs", handlers.BlogHandler).Methods("GET")

	// Server
	log.Fatal(http.ListenAndServe(":8081", r))
}
