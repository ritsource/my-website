package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

var api string

func init() {
	api = os.Getenv("API_URI")
}

// Blog
type Blog struct {
	ID          string `json:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HTML        string `json:"html"`
	Markdown    string `json:"markdown"`
	DocType     string `json:"doc_type"`
	IsPublic    bool   `json:"is_public"`
	IsDeleted   bool   `json:"is_deleted"`
}

// BlogHandler ...
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	// Get Public Data from API
	resp, err := http.Get("http://" + api + "/public/blog/all")
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	defer resp.Body.Close()

	// Reading Response Body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	// Unmarshaling Body Data
	var data []Blog

	err = json.Unmarshal(b, &data)
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/blogs.html",
		"static/partials/header.html",
		"static/partials/blogs-item.html",
	)
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}

	// Executing Template
	err = t.Execute(w, data)
	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
