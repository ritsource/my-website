package handlers

import (
	"net/http"
	"text/template"
)

// BlogHandler ...
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/pages/blogs.html", "static/partials/header.html")
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}

	err = t.Execute(w, []string{})
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
