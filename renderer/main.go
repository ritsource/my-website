package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	// Router (Gorilla Mux)
	r := mux.NewRouter()

	// Dummy handler
	r.HandleFunc("/", indexHandler).Methods("GET")

	// Server
	log.Fatal(http.ListenAndServe(":8081", r))
}

// WriteError Admin Un-Authenticated on Response
func WriteError(w http.ResponseWriter, status int, err error, msg string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\"}"))
	fmt.Println("Error:", err.Error())
}

// Dummy Index Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html", "static/header.html")
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}

	err = t.Execute(w, []string{})
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
