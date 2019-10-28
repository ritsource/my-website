package renderers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/ritcrap/my-website/server/db"
	"github.com/ritcrap/my-website/server/raw"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BlogHandler renders a single blog
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	// if just `ritwiksaha.com/blog/`, then redirecting to `/blogs` route
	if paths[2] == "" {
		http.Redirect(w, r, "/blogs", http.StatusSeeOther)
		return
	}

	idstr := paths[2]

	// reading from blog document database
	var b db.Blog
	err := b.Read(bson.M{"id_str": idstr, "is_deleted": false, "is_public": true}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		renderErr(w, 404, fmt.Sprintf("Blog \"%v\" Not Found", idstr))
	case nil:
		// everything's fine
	default:
		// some internal error
		writeErr(w, 500, err)
		return
	}

	if b.IsSeries {
		http.Redirect(w, r, fmt.Sprintf("/thread/%v", idstr), http.StatusSeeOther)
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
	// also, in raw.GetDocument arguements index will always
	// be "0" as there's just 1 document
	doc, err := raw.GetDocument(b.ID.Hex(), src, b.DocType, 0)
	if err != nil {
		renderErr(w, 500, "Couldn't Read Document")
		return
	}

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if b.DocType == db.DocTypeMD {
		unsafe = blackfriday.MarkdownCommon(doc) // generating HTML from Markdown
	} else {
		unsafe = doc
	}

	// Document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-doc.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data db.Blog
		HTML string
	}{
		Data: b,
		HTML: string(html),
	})

	if err != nil {
		writeErr(w, 500, err)
	}

}

// ThreadHandler renders all the blogs
func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	// if just `ritwiksaha.com/blog/`, then redirecting to `/blogs` route
	if paths[2] == "" {
		renderErr(w, 400, "No Thread-ID Provided")
		return
	}

	idstr := paths[2]

	// reading requested index from URL
	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	if err != nil || index < 0 {
		renderErr(w, 400, "Invalid Index")
		return
	}

	// reading from blog/thread document database
	var b db.Blog
	err = b.Read(bson.M{"id_str": idstr, "is_deleted": false, "is_public": true, "is_series": true}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		renderErr(w, 404, fmt.Sprintf("Thread \"%v\" Not Found", idstr))
	case nil:
		// everything's fine
	default:
		// some internal error
		writeErr(w, 500, err)
		return
	}

	// If thread doesn't include any subblog
	if len(b.SubBlogs) == 0 {
		renderErr(w, http.StatusNoContent, "Sorry, Empty Thread")
		return
	}

	// If index overflow, then redirect to index 0
	if index+1 > len(b.SubBlogs) {
		http.Redirect(w, r, fmt.Sprintf("/thread/%v?index=0", idstr), http.StatusSeeOther)
		return
	}

	// remote src string
	var src string
	switch b.SubBlogs[index].DocType {
	case db.DocTypeMD:
		src = b.SubBlogs[index].Markdown
	case db.DocTypeHTML:
		src = b.SubBlogs[index].HTML
	}

	// reading the document, for index there's some additional characters in the end
	doc, err := raw.GetDocument(b.ID.Hex(), src, b.DocType, index)
	if err != nil {
		renderErr(w, 500, "Couldn't Read Document")
		return
	}

	// parsing document data (html or markdown) into HTML (unsafe)
	var unsafe []byte
	if b.SubBlogs[index].DocType == db.DocTypeMD {
		unsafe = blackfriday.MarkdownCommon(doc) // generating HTML from Markdown
	} else {
		unsafe = doc
	}

	// Document HTML
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/each-thread.html",
		"static/partials/thread-nav.html",
		"static/partials/header.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		renderErr(w, 500, "Internal Server Error")
		return
	}

	var prevURL string
	var nextURL string

	if index == 0 {
		prevURL = ""
	} else {
		prevURL = fmt.Sprintf("/thread/%v?index=%v", b.IDStr, strconv.Itoa(index-1))
	}

	if index+1 < len(b.SubBlogs) {
		nextURL = fmt.Sprintf("/thread/%v?index=%v", b.IDStr, strconv.Itoa(index+1))
	} else {
		nextURL = ""
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data       db.Blog
		SubBlog    db.SubBlog
		Index      int
		PrevSubURL string
		NextSubURL string
		HTML       string
	}{
		Data:       b,
		SubBlog:    b.SubBlogs[index],
		Index:      index,
		PrevSubURL: prevURL,
		NextSubURL: nextURL,
		HTML:       fmt.Sprintf("%s\n", html),
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
		sel = bson.M{"is_deleted": false, "is_public": true, "is_technical": true}
	} else {
		sel = bson.M{"is_deleted": false, "is_public": true, "is_technical": false}
	}

	// querying data from database
	var bs db.Blogs
	err := bs.ReadAll(sel, bson.M{})
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
