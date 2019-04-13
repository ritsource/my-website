package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// Public Routes - Public routes are publically accessable (as the name suggests)
// it doesn't require admin authentication

// Blog Routes

// PubReadBlogs -
func PubReadBlogs(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlogs models.Blogs

	// Read blog
	mBlogs, err = mBlogs.Read(bson.M{
		"is_deleted": false,
		"is_public":  true,
	})
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

// PubReadBlogByID - Read Public blog by ID
func PubReadBlogByID(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlog models.Blog

	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Read blog
	mBlog, err = mBlog.ReadSingle(bson.M{
		"_id":        bson.ObjectIdHex(bIDStr),
		"is_deleted": false,
		"is_public":  true,
	})
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

// PubGetBlogDocument - This Handler Queries Blog from the database first, then checks
// if file exists in the cache or not, (in the "/satic" folder), (K8s PV in Prod),
// If file exists then serves it else redirects the user to the source
// saved in the mongo document, also if file doesn't exist the it downloads
// the file in the static folder
func PubGetBlogDocument(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlog models.Blog // Blog struct (models.Blog)

	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Reading the Blog document from database
	mBlog, err = mBlog.ReadSingle(bson.M{
		"_id":        bson.ObjectIdHex(bIDStr),
		"is_deleted": false,
		"is_public":  true,
	})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Defining filename for in cache folder
	var fileName string
	if mBlog.DocType == "markdown" {
		fileName = bIDStr + ".md"
	} else {
		fileName = bIDStr + ".html"
	}

	// Checking if requested file Exists or Not
	if _, err := os.Stat("./cache/" + fileName); err == nil {
		// Redirect to static route if file exist
		http.Redirect(w, r, "/static/"+fileName, http.StatusTemporaryRedirect)
		return
	}

	// If file doesn't exist in cache
	fmt.Println("File Doesn't Exist!", fileName)

	// Get file source (markdown or html)
	var srcFilePath string
	if mBlog.DocType == "markdown" {
		srcFilePath = mBlog.Markdown
	} else {
		srcFilePath = mBlog.HTML
	}

	// Redirecting to the source file
	http.Redirect(w, r, srcFilePath, http.StatusTemporaryRedirect)

	// Downloading the file in cache
	DownloadFile("./cache/"+fileName, srcFilePath)
}

// Project Routes

// PubReadProjects -
func PubReadProjects(w http.ResponseWriter, r *http.Request) {
	var err error
	var mProjects models.Projects

	// Read Project
	mProjects, err = mProjects.Read(bson.M{
		"is_deleted": false,
		"is_public":  true,
	})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mProjects)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// PubReadProjectByID - Read Public Project By ID
func PubReadProjectByID(w http.ResponseWriter, r *http.Request) {
	var err error
	var mProject models.Project

	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Read project
	mProject, err = mProject.ReadSingle(bson.M{
		"_id":        bson.ObjectIdHex(pIDStr),
		"is_deleted": false,
		"is_public":  true,
	})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mProject)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// PubGetProjectDocument - This Handler Queries Project from the database first, then checks
// if file exists or not in the "/satic" folder, (K8s PV in Prod),
// If file exists then serves it else redirects the user to the source
// saved in the mongo document, also if file doesn't exist the it downloads
// the file in the static folder
func PubGetProjectDocument(w http.ResponseWriter, r *http.Request) {
	var err error
	var mProject models.Project // Project struct

	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Reading Project from the database
	mProject, err = mProject.ReadSingle(bson.M{
		"_id":        bson.ObjectIdHex(pIDStr),
		"is_deleted": false,
		"is_public":  true,
	})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Defining cahce filename
	var fileName string
	if mProject.DocType == "markdown" {
		fileName = pIDStr + ".md"
	} else {
		fileName = pIDStr + ".html"
	}

	// Checking if requested file Exists or Not
	if _, err := os.Stat("./cache/" + fileName); err == nil {
		// if exist then redirect to the static route
		http.Redirect(w, r, "/static/"+fileName, http.StatusTemporaryRedirect)
		return
	}

	// If file doesn't exist in cache
	fmt.Println("File Doesn't Exist!", fileName)

	// File Source
	var srcFilePath string
	if mProject.DocType == "markdown" {
		srcFilePath = mProject.Markdown
	} else {
		srcFilePath = mProject.HTML
	}

	// Redirecting to the main source file
	http.Redirect(w, r, srcFilePath, http.StatusTemporaryRedirect)

	// Downloading the file
	DownloadFile("./cache/"+fileName, srcFilePath)
}

// DownloadFile - Used for saving files in the cache
// Downloads a document from web
// "path" for local path (where to save), "src" for the url
func DownloadFile(path string, src string) {
	// Get response from teh url
	resp, err := http.Get(src)
	if err != nil {
		fmt.Printf("Unable to read source file on web: %v", err)
	}

	defer resp.Body.Close()

	// Creating a new file on given path
	out, err := os.Create(path)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	defer out.Close()

	// Writing the body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
