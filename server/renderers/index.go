package renderers

import (
	"net/http"
	"text/template"
)

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	t, err := template.ParseFiles(
		"static/pages/index.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		writeErr(w, 500, err)
	}

	err = t.Execute(w, []string{})
	if err != nil {
		writeErr(w, 500, err)
	}
}

// ResumeHandler .
func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/pages/resume.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		writeErr(w, 500, err)
	}

	err = t.Execute(w, []string{})
	if err != nil {
		writeErr(w, 500, err)
	}
}
