package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func writeErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)                                     // Status Code
	w.Header().Set("Content-Type", "application/json")        // Response Type - JSON
	w.Write([]byte("{\"message\": \"" + err.Error() + "\"}")) // Error Message

	logrus.Warnf("%v\n", err)
}

// writeJSON writes json to the http.ResponseWriter
func writeJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
