package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

// PreviewHandler - This handler renders document for any given document source
// for example http://localhost:8081/preview?type=markdown&src=https://somewhere/README.md will
// render a document with https://somewhere/README.md (src) file
// Useful for Previewing Private documents
// Example request ...
// http://localhost:8081/preview?type=markdown&src=https://raw.githubusercontent.com/ritwik310/markdown-editor/master/README.md
func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query() // Query String Values

	// Check if required query keys exist or not
	if len(query["type"]) < 1 || len(query["src"]) < 1 {
		WriteError(w, 500, errors.New("Not Enough Query Values"), "Not Enough Query Values")
		return
	}

	// Document Source
	src := query["src"][0]
	docType := query["type"][0] // Document type (HTML or Markdown)

	c := make(chan []byte) // Channel for Data Fetching

	// Get Public Data from API
	go FetchData(src, c)

	// Fetched byte Data
	b := <-c

	// Check Error
	if b == nil {
		WriteError(w, 500, errors.New("Couldn't Fetch Data"), "Couldn't Fetch Data")
		return
	}

	// Unsafe HTML (From Doc)
	var unsafe []byte

	if docType == "markdown" {
		// Generating HTML from Markdown
		unsafe = blackfriday.Run(b)
	} else if docType == "html" {
		unsafe = b
	} else {
		WriteError(w, 500, errors.New("Doctype not defined"), "Doctype not defined")
		return
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

	// Dummy Blog Data
	data := Blog{
		ID:          "PreviewID",
		Title:       "Preview Title",
		Description: "Preview Description",
		HTML:        src,
		Markdown:    src,
		DocType:     docType,
		IsPublic:    true,
		IsDeleted:   false,
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data    Blog
		HTML    string
		Project bool
	}{
		Data:    data,
		HTML:    fmt.Sprintf("%s\n", html),
		Project: false, // False because it can be anything (blog or project)
	})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
