package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/russross/blackfriday.v2"
)

// API domain
var API string

func init() {
	isDev := os.Getenv("DEV_MODE") == "true"

	if isDev {
		API = os.Getenv("API_URL")
	} else {
		// Docker networking will  take care of that (docker-compose, or aws)
		API = "http://api:8080"
	}
}

// Blog type - In data fetched from API
type Blog struct {
	ID              string    `json:"_id,omitempty"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DescriptionLink string    `json:"description_link"`
	Author          string    `json:"author"`
	FormattedDate   string    `json:"formatted_date"`
	HTML            string    `json:"html"`
	Markdown        string    `json:"markdown"`
	DocType         string    `json:"doc_type"`
	Thumbnail       string    `json:"thumbnail"`
	IsTechnical     bool      `json:"is_technical"`
	IsPublic        bool      `json:"is_public"`
	IsDeleted       bool      `json:"is_deleted"`
	IsSeries        bool      `bson:"is_series" json:"is_series"`
	SubBlogs        []SubBlog `bson:"sub_blogs" json:"sub_blogs"`
}

// SubBlog - Blog model type
type SubBlog struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title         string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	FormattedDate string        `bson:"formatted_date" json:"formatted_date"`
	HTML          string        `bson:"html" json:"html"`
	Markdown      string        `bson:"markdown" json:"markdown"`
	DocType       string        `bson:"doc_type" json:"doc_type"`
}

// FetchDataAsync - Fetches Data from the API
func FetchDataAsync(url string, c chan []byte) {
	// Get Public Data from API
	resp, err := http.Get(url)
	if err != nil {
		c <- nil
		return
	}

	defer resp.Body.Close()

	// Reading Response Body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- nil
		return
	}

	c <- b
}

// FetchDataSync ...
func FetchDataSync(url string) ([]byte, error) {
	// Get Public Data from API
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Reading Response Body
	return ioutil.ReadAll(resp.Body)
}

// EachThreadHandler ..
func EachThreadHandler(w http.ResponseWriter, r *http.Request) {
	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String
	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	if err != nil || index < 0 {
		RenderError(w, 400, "Invalid Index")
		return
	}

	// Get Public Data from API
	b1, err := FetchDataSync(API + "/api/public/blog/" + bIDStr)
	if err != nil {
		RenderError(w, 404, "Blog Not Found")
		return
	}

	// Unmarshaling Body Data
	var data Blog
	err = json.Unmarshal(b1, &data)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
		return
	}

	if !data.IsSeries {
		RenderError(w, 400, "Not A Thread")
		return
	}

	if len(data.SubBlogs) == 0 {
		RenderError(w, 400, "Empty Thread")
		return
	}

	if index+1 < len(data.SubBlogs) {
		index = 0
	}

	var docSrc string
	isMd := data.SubBlogs[index].DocType == "markdown"

	if isMd {
		docSrc = data.SubBlogs[index].Markdown
	} else {
		docSrc = data.SubBlogs[index].HTML
	}

	b2, err := FetchDataSync(docSrc)
	if err != nil {
		RenderError(w, 404, "Document Not Found")
		return
	}

	// Unsafe HTML (From Doc)
	var unsafe []byte
	if isMd {
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
		RenderError(w, 500, "Internal Server Error")
		return
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data    Blog
		HTML    string
		Project bool
	}{
		Data:    data,
		HTML:    fmt.Sprintf("%s\n", html),
		Project: false,
	})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}

// EachBlogHandler - Fetches single Blog data and Document for that Blog,
// Renders document (Html or Markdown) inside HTML template
func EachBlogHandler(w http.ResponseWriter, r *http.Request) {
	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	c1 := make(chan []byte) // Channel for Data Fetching
	c2 := make(chan []byte) // Channel for Document Fetching

	// Get Public Data from API
	go FetchDataAsync(API+"/api/public/blog/"+bIDStr, c1)

	// Get Public Data from API
	go FetchDataAsync(API+"/api/public/blog/doc/"+bIDStr, c2)

	b1 := <-c1 // Blog data 1 ([]byte)
	b2 := <-c2 // Document data 2 ([]byte)

	// Check Error
	if b1 == nil || b2 == nil {
		RenderError(w, 404, "Blog Not Found")
		return
	}

	// Unmarshaling Body Data
	var data Blog

	err := json.Unmarshal(b1, &data)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
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
		RenderError(w, 500, "Internal Server Error")
		return
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data    Blog
		HTML    string
		Project bool
	}{
		Data:    data,
		HTML:    fmt.Sprintf("%s\n", html),
		Project: false,
	})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}

// BlogsHandler - Handler for All Blogs Page
func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	// channel
	c := make(chan []byte)
	// var blogsroute string
	var tech bool
	var cookie http.Cookie

	// Checking if technical only
	qTech := r.URL.Query()["tech"]

	// Checking if query string exist or not
	if len(qTech) > 0 {
		if qTech[0] == "true" {
			tech = true
			cookie = http.Cookie{Name: "tech", Value: "true", Path: "/", Expires: time.Now().AddDate(0, 1, 0), MaxAge: 86400} // Cookie
		} else if qTech[0] == "false" {
			cookie = http.Cookie{Name: "tech", Value: "false", Path: "/", Expires: time.Now().AddDate(0, 1, 0), MaxAge: 86400} // Cookie
		}
	} else {
		// If nothing in query string read read from cookie
		var tCuki, err = r.Cookie("tech")
		// If value exists and its "true", then make tech true
		if err == nil && tCuki.Value == "true" {
			tech = true
		}
	}

	// Get Public Data from API
	if tech {
		go FetchDataAsync(API+"/api/public/blog/all?tech=true", c)
	} else {
		go FetchDataAsync(API+"/api/public/blog/all?tech=false", c)
	}

	// Unmarshaling Body Data
	var data []Blog

	err := json.Unmarshal(<-c, &data)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
		return
	}

	// Parsing templates
	t, err := template.ParseFiles(
		"static/pages/blogs.html",
		"static/partials/header.html",
		"static/partials/blogs-item.html",
		"static/partials/social-btns.html",
	)
	if err != nil {
		RenderError(w, 500, "Internal Server Error")
	}

	http.SetCookie(w, &cookie)

	// Executing Template
	err = t.Execute(w, struct {
		Tech bool
		Data []Blog
	}{
		Tech: tech,
		Data: data,
	})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}
