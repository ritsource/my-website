package renderers

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/raw"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/russross/blackfriday.v2"
)

// BlogHandler renders a single blog
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query()["id"]

	paths := strings.Split(r.URL.Path, "/")

	if paths[2] == "" {
		http.Redirect(w, r, "/blogs", http.StatusPermanentRedirect)
		return
	}

	id := paths[2]

	// reading from blog document database
	var b db.Blog
	err := b.Read(bson.M{"id_str": id, "is_deleted": false, "is_public": true}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		renderErr(w, 404, fmt.Sprintf("Blog \"%v\" Not Found", id))
	case nil:
		// everything's fine
	default:
		// some internal error
		writeErr(w, 500, err)
		return
	}

	// remote src string
	var src string
	switch b.DocType {
	case db.DocTypeMD:
		src = b.Markdown
	case db.DocTypeHTML:
		src = b.HTML
	}

	// reading the document, (`GetDocument` handles caching)
	doc, err := raw.GetDocument(b.ID.String(), src, b.DocType)
	if err != nil {
		renderErr(w, 500, "Couldn't Read Document")
		return
	}

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if b.DocType == db.DocTypeMD {
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
		Data    db.Blog
		HTML    string
		Project bool
	}{
		Data:    b,
		HTML:    string(html),
		Project: false,
	})

	if err != nil {
		writeErr(w, 500, err)
	}

}

// BlogsHandler renders all the blogs
func BlogsHandler(w http.ResponseWriter, r *http.Request) {

	var tech bool
	var cookie http.Cookie

	qTech := r.URL.Query()["tech"]

	// if tech is included in the query string then render
	// only technical content `"is_technical": true`, and
	// set the default query type as tech in the cookie.
	// else, render non-technical content and set default
	// as non-tech in cookie. if there's nothing in the
	// query string, then render as the default in cookie
	if len(qTech) > 0 {
		// if query string exist
		if qTech[0] == "true" {
			// if `?tech=true` in query string
			tech = true
			cookie = http.Cookie{Name: "tech", Value: "true", Path: "/", Expires: time.Now().AddDate(0, 1, 0), MaxAge: 86400} // Cookie
		} else if qTech[0] == "false" {
			// if `?tech=false` in query str
			cookie = http.Cookie{Name: "tech", Value: "false", Path: "/", Expires: time.Now().AddDate(0, 1, 0), MaxAge: 86400} // Cookie
		}
	} else {
		// If nothing in query string, read from cookie
		var tCookie, err = r.Cookie("tech")
		// If value exists and its "true", then make tech true
		if err == nil && tCookie.Value == "true" {
			tech = true
		}
	}

	// constructing the selector for query from mongodb
	var sel bson.M
	if tech {
		sel = bson.M{"is_deleted": false, "is_public": true, "is_technical": false}
	} else {
		sel = bson.M{"is_deleted": false, "is_public": true, "is_technical": true}
	}

	// querying data from database
	var bs db.Blogs
	err := bs.Read(sel, bson.M{})
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
	}

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/blogs.html",
		"static/partials/header.html",
		"static/partials/blogs-item.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
	}

	// setting cookie appropriately
	http.SetCookie(w, &cookie)

	// Executing Template
	err = t.Execute(w, struct {
		Tech bool
		Data db.Blogs
	}{
		Tech: tech,
		Data: bs,
	})

	if err != nil {
		writeErr(w, 500, err)
	}

}
