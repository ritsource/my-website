package renderers

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/raw"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/russross/blackfriday.v2"
)

// ProjectHandler renders a single project description
func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	// if just `ritwiksaha.com/project/`, then redirecting to `/projects` route
	if paths[2] == "" {
		http.Redirect(w, r, "/projects", http.StatusSeeOther)
		return
	}

	idstr := paths[2]

	// reading from project document database
	var p db.Project
	err := p.Read(bson.M{"id_str": idstr, "is_deleted": false, "is_public": true}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		renderErr(w, 404, fmt.Sprintf("Project \"%v\" Not Found", idstr))
	case nil:
		// everything's fine
	default:
		// some internal error
		writeErr(w, 500, err)
		return
	}

	// remote src string
	var src string
	switch p.DocType {
	case db.DocTypeMD:
		src = p.Markdown
	case db.DocTypeHTML:
		src = p.HTML
	}

	// reading the document, (`GetDocument` handles caching)
	// also, in raw.GetDocument arguements index will always
	// be "0" as there's just 1 document
	doc, err := raw.GetDocument(p.ID.Hex(), src, p.DocType, 0)
	if err != nil {
		renderErr(w, 500, "Couldn't Read Document")
		return
	}

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if p.DocType == db.DocTypeMD {
		unsafe = blackfriday.Run(doc) // generating HTML from Markdown
	} else {
		unsafe = doc
	}

	// Document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-doc.html",
		"static/partials/header.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data    db.Project
		HTML    string
		Project bool
	}{
		Data:    p,
		HTML:    string(html),
		Project: true,
	})

	if err != nil {
		writeErr(w, 500, err)
	}

}

// ProjectsHandler renders all the projects
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {

	// querying data from database
	var ps db.Projects
	err := ps.Read(bson.M{"is_deleted": false, "is_public": true}, bson.M{})
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
	}

	// parsing templates
	t, err := template.ParseFiles(
		"static/pages/projects.html",
		"static/partials/header.html",
		"static/partials/projects-item.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
	}

	// executing Template
	err = t.Execute(w, ps)
	if err != nil {
		writeErr(w, 500, err)
	}

}
