package renderers

import (
	"net/http"
	"text/template"

	"github.com/sirupsen/logrus"
)

func writeErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)                                     // Status Code
	w.Header().Set("Content-Type", "application/json")        // Response Type - JSON
	w.Write([]byte("{\"message\": \"" + err.Error() + "\"}")) // Error Message

	logrus.Warnf("%v\n", err)
}

// renderErr renders out HTML showing the Error
func renderErr(w http.ResponseWriter, status int, message string) {
	t, err := template.ParseFiles(
		"static/pages/error.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		writeErr(w, 500, err)
	}

	w.WriteHeader(status)

	err = t.Execute(w, struct {
		Status  int
		Message string
	}{Status: status, Message: message})

	if err != nil {
		writeErr(w, 500, err)
	}
}
