package renderers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/ritcrap/my-website/server/db"
	"github.com/russross/blackfriday"
)

// PreviewHandler renders document for any given document source
func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// check if required query keys exist or not
	if len(query["type"]) < 1 || len(query["src"]) < 1 {
		renderErr(w, http.StatusBadRequest, "Not Enough Query Values")
		return
	}

	src := query["src"][0]
	var doctype int8
	switch query["type"][0] {
	case "markdown":
		doctype = db.DocTypeMD
	case "html":
		doctype = db.DocTypeHTML
	default:
		renderErr(w, http.StatusBadRequest, "Invalid Document Type")
		return
	}

	// reading data from remote source
	resp, err := http.Get(src)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}
	defer resp.Body.Close()

	doc, err := ioutil.ReadAll(resp.Body)

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if doctype == db.DocTypeMD {
		unsafe = blackfriday.MarkdownCommon(doc) // generating HTML from Markdown
	} else {
		unsafe = doc
	}

	// Document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-doc.html",
		"static/partials/header.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}

	// dummy object data
	data := db.Blog{
		ID:            "PreviewID",
		Title:         "Preview Title",
		Description:   "Preview Description",
		Author:        "Ritwik Saha",
		FormattedDate: "April 2, 2019",
		HTML:          src,
		Markdown:      src,
		DocType:       doctype,
		IsPublic:      true,
		IsDeleted:     false,
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data db.Blog
		HTML string
	}{
		Data: data,
		HTML: fmt.Sprintf("%s", html),
	})

	if err != nil {
		writeErr(w, 500, err)
	}
}
