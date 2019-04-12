package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// WriteError Admin Un-Authenticated on Response
func WriteError(w http.ResponseWriter, status int, err error, msg string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\"}"))
	fmt.Println("Error:", err.Error())
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html", "static/header.html")
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}

	err = t.Execute(w, []string{})
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
