package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

var api string

func init() {
	api = os.Getenv("API_URI")
}

// Blog ...
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

// FetchData - Fetches Data
func FetchData(url string, c chan []byte) {
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

// EachBlogHandler ...
func EachBlogHandler(w http.ResponseWriter, r *http.Request) {
	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	c1 := make(chan []byte) // Channel for Data Fetching
	c2 := make(chan []byte) // Channel for Document Fetching

	// Get Public Data from API
	go FetchData("http://"+api+"/public/blog/"+bIDStr, c1)

	// Get Public Data from API
	go FetchData("http://"+api+"/public/blog/doc/"+bIDStr, c2)

	b1 := <-c1 // Blog data 1 ([]byte)
	b2 := <-c2 // Document data 2 ([]byte)

	if b1 == nil || b2 == nil {
		WriteError(w, 500, errors.New("Couldn't Fetch Data"), "Couldn't Fetch Data")
		return
	}

	// Unmarshaling Body Data
	var data Blog

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
		"static/pages/each-blog.html",
		"static/partials/header.html",
		"static/partials/blogs-item.html",
	)
	if err != nil {
		WriteError(w, 500, err, err.Error())
		return
	}

	// Executing Template
	err = t.Execute(w, struct {
		Data Blog
		HTML string
	}{
		Data: data,
		HTML: fmt.Sprintf("%s\n", html),
	})

	if err != nil {
		WriteError(w, 500, err, err.Error())
	}
}

// BlogsHandler ...
func BlogsHandler(w http.ResponseWriter, r *http.Request) {
	// channel
	c := make(chan []byte)

	// Get Public Data from API
	go FetchData("http://"+api+"/public/blog/all", c)

	// Unmarshaling Body Data
	var data []Blog

	err := json.Unmarshal(<-c, &data)
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
