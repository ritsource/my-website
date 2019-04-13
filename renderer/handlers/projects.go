package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

// Project type - In data fetched from API
type Project struct {
	ID          string `json:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HTML        string `json:"html"`
	Markdown    string `json:"markdown"`
	DocType     string `json:"doc_type"`
	Link        string `json:"link"`
	IsPublic    bool   `json:"is_public"`
	IsDeleted   bool   `json:"is_deleted"`
}

// EachProjectHandler - Fetches single Project data and Document for that Project,
// Renders document (Html or Markdown) inside HTML template
func EachProjectHandler(w http.ResponseWriter, r *http.Request) {
	pIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	c1 := make(chan []byte) // Channel for Data Fetching
	c2 := make(chan []byte) // Channel for Document Fetching

	// Get Public Data from API
	go FetchData("http://"+API+"/public/project/"+pIDStr, c1)

	// Get Public Data from API
	go FetchData("http://"+API+"/public/project/doc/"+pIDStr, c2)

	b1 := <-c1 // Blog data 1 ([]byte)
	b2 := <-c2 // Document data 2 ([]byte)

	// Check Error
	if b1 == nil || b2 == nil {
		WriteError(w, 500, errors.New("Couldn't Fetch Data"), "Couldn't Fetch Data")
		return
	}

	// Unmarshaling Body Data
	var data Project

	err := json.Unmarshal(b1, &data)
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	// Unsafe HTML (From Doc)
	var unsafe []byte

	if data.DocType == "markdown" {
		// Generating HTML from Markdown
		unsafe = blackfriday.Run(b2)
	} else {
		unsafe = b2
	}

	// Document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-doc.html",
		"static/partials/header.html",
	)
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data    Project
		HTML    string
		Project bool
	}{
		Data:    data,
		HTML:    fmt.Sprintf("%s\n", html),
		Project: true,
		// Project: data.Link != "",
	})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}

// ProjectsHandler - Handler for All Projects Page
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	// channel
	c := make(chan []byte)

	// Get Public Data from API
	go FetchData("http://"+API+"/public/project/all", c)

	// Unmarshaling Body Data
	var data []Project

	err := json.Unmarshal(<-c, &data)
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/projects.html",
		"static/partials/header.html",
		"static/partials/projects-item.html",
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
