package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ritwik310/my-website/server/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateBlog creates a new blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading JSON body
	var b db.Blog

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// creation time
	b.CreatedAt = int32(time.Now().Unix())

	// insert new document
	err = b.Create()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// redirecting to `/blogs` route handler
	http.Redirect(w, r, "/api/private/blogs", http.StatusTemporaryRedirect) // 302 - POST to GET
}

// ReadBlog reads a single blog
func ReadBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// reading document
	var b db.Blog
	err := b.Read(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{})
	switch err {
	case mgo.ErrNotFound:
		writeErr(w, 404, err)
	case nil:
		// everything's fine
	default:
		writeErr(w, 500, err)
		return
	}

	// writing json data to the client
	writeJSON(w, b)
}

// ReadBlogs reads all blogs
func ReadBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	var bs db.Blogs
	err := bs.ReadAll(bson.M{}, bson.M{})
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	writeJSON(w, bs)
}

// EditBlog dits a blog by `_id`
func EditBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// reading request body
	var body map[string]interface{} // because cannot use b (type db.Blog) as type bson.M in argument to bl.Update
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// editing document
	var b db.Blog
	err = b.Update(bson.M{"_id": bson.ObjectIdHex(id)}, body) // Update Document in Database
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, b) // Write Data
}

// DeleteBlog - Deletes a Blog (Not Permanently)
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// editing document to `is_deleted: true`
	var b db.Blog
	err := b.Delete(bson.ObjectIdHex(id))
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// writing the updated data
	writeJSON(w, b)
}

// DeleteBlogPrem - Deletes a blog Permanently
func DeleteBlogPrem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no blog-id provided"))
		return
	}

	// deleting document (permanently)
	var b db.Blog
	b.ID = bson.ObjectIdHex(id)
	err := b.DeletePermanent()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
