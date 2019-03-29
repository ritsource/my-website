package routes

import (
	"encoding/json"
	"os"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateBlog ...
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var body models.Blog // to save Blog JSON body
	var err error

	// Decoding request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		WriteError(w, 422, err, "Unable to read request body")
		return
	}

	// Inserting Document
	nBlog, err := body.Create()
	if err != nil {
		WriteError(w, 422, err, "Failed to insert new document")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nBlog)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadOneBlog - ...
func ReadOneBlog(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlog models.Blog

	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Read blog
	mBlog, err = mBlog.ReadSingle(bson.M{"_id": bson.ObjectIdHex(bIDStr)})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mBlog)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// GetBlogDocument - 
// This Handler Queries Blog from teh database first, then checks
// if file exists or not in the "/satic" folder, (K8s PV in Prod),
// If file exists then serves it else redirects the user to the source
// saved in the mongo document, also if file doesn't exist the it downloads
// the file in the static folder
func GetBlogDocument(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlog models.Blog

	pIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Read Blog
	mBlog, err = mBlog.ReadSingle(bson.M{"_id": bson.ObjectIdHex(pIDStr)})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Generating Local file Name
	var fileName string
	if mBlog.DocType == "markdown" {
		fileName = pIDStr + ".md"
	} else {
		fileName = pIDStr + ".html"
	}
	
	// Checking if requested file Exists or Not
	if _, err := os.Stat("./static/" + fileName); err == nil {
		http.Redirect(w, r, "/static/" + fileName, http.StatusTemporaryRedirect)
		return
	} else {
		fmt.Println("File Doesn't Exist!", fileName)
	}

	// File Source
	var srcFilePath string
	if mBlog.DocType == "markdown" {
		srcFilePath = mBlog.Markdown
	} else {
		srcFilePath = mBlog.HTML
	}

	// Redirecting to the source file
	http.Redirect(w, r, srcFilePath, http.StatusTemporaryRedirect)
	
	// Downloading the file
	DownloadFile("./static/" + fileName, srcFilePath)
}

// ReadBlogs - read all blogs, both Public and Private
func ReadBlogs(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlogs models.Blogs

	// Read blog
	mBlogs, err = mBlogs.Read(bson.M{})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mBlogs)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// EditBlog - ...
func EditBlog(w http.ResponseWriter, r *http.Request) {
	var err error
	var nBlog models.Blog

	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Decoding request body
	var body map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		WriteError(w, 422, err, "Unable to read request body")
		return
	}

	fmt.Printf("body %+v\n", body)

	// Update Blog Document
	nBlog, err = nBlog.Update(
		bson.M{"_id": bson.ObjectIdHex(bIDStr)},
		body,
	)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nBlog)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// DeleteBlog - ...
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	var err error
	var dBlog models.Blog

	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Read blog
	_, err = dBlog.Delete(bson.ObjectIdHex(bIDStr))
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
