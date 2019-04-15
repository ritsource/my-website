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

// RenderError - Renderes out HTML showing the Error
func RenderError(w http.ResponseWriter, status int16, message string) {
	t, err := template.ParseFiles(
		"static/pages/error.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}

	err = t.Execute(w, struct {
		Status  int16
		Message string
	}{Status: status, Message: message})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/pages/index.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}

	err = t.Execute(w, []string{})
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
