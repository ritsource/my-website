package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// DownloadFile - Used for saving files in the cache
// Downloads a document from web
// "path" for local path (where to save), "src" for the url
func DownloadFile(path string, src string) error {
	// Get response from teh url
	resp, err := http.Get(src)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Creating a new file on given path
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Writing the body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// PubGetBlogDoc - This Handler Queries Blog from the database first, then checks
// if file exists in the cache or not, (in the "/satic" folder), (K8s PV in Prod),
// If file exists then serves it else redirects the user to the source
// saved in the mongo document, also if file doesn't exist the it downloads
// the file in the static folder
func PubGetBlogDoc(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog
	bIDStr := mux.Vars(r)["id"]

	bl, err := bl.Read(bson.M{"_id": bson.ObjectIdHex(bIDStr), "is_deleted": false, "is_public": true}) // Reading Data
	HandleErr(w, 422, err)

	// Defining filename for in cache folder
	var fileName string
	if bl.DocType == "markdown" {
		fileName = bIDStr + ".md"
	} else {
		fileName = bIDStr + ".html"
	}

	// Checking if requested file Exists or Not
	if _, err := os.Stat(path.Join(".", "cache", fileName)); err == nil {
		// Redirect to static route if file exist
		http.Redirect(w, r, "/api/static/"+fileName, http.StatusTemporaryRedirect)
		return
	}

	// Get file source (markdown or html)
	var srcFilePath string
	if bl.DocType == "markdown" {
		srcFilePath = bl.Markdown
	} else {
		srcFilePath = bl.HTML
	}

	http.Redirect(w, r, srcFilePath, http.StatusTemporaryRedirect) // Redirecting to the source file

	err = DownloadFile(path.Join(".", "cache", fileName), srcFilePath) // Downloading the file in cache
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// PubGetProjectDoc - Same a the above for Project Documents
func PubGetProjectDoc(w http.ResponseWriter, r *http.Request) {
	var pr models.Project
	pIDStr := mux.Vars(r)["id"]

	pr, err := pr.Read(bson.M{"_id": bson.ObjectIdHex(pIDStr), "is_deleted": false, "is_public": true})
	HandleErr(w, 422, err)

	// Defining cahce filename
	var fileName string
	if pr.DocType == "markdown" {
		fileName = pIDStr + ".md"
	} else {
		fileName = pIDStr + ".html"
	}

	// Checking if requested file Exists or Not
	if _, err := os.Stat(path.Join(".", "cache", fileName)); err == nil {
		// if exist then redirect to the static route
		http.Redirect(w, r, "/api/static/"+fileName, http.StatusTemporaryRedirect)
		return
	}

	// File Source
	var srcFilePath string
	if pr.DocType == "markdown" {
		srcFilePath = pr.Markdown
	} else {
		srcFilePath = pr.HTML
	}

	http.Redirect(w, r, srcFilePath, http.StatusTemporaryRedirect) // Redirecting to the source file

	err = DownloadFile(path.Join(".", "cache", fileName), srcFilePath) // Downloading the file in cache
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// ClearDocCache - Clears cache files
func ClearDocCache(w http.ResponseWriter, r *http.Request) {
	iDStr := mux.Vars(r)["id"]

	mdFile := iDStr + ".md"
	htmlFile := iDStr + ".html"

	// If exists then delete the markdown file
	if _, err := os.Stat(path.Join(".", "cache", mdFile)); err == nil {
		err = os.Remove(path.Join(".", "cache", mdFile))
		HandleErr(w, 500, err)
	}

	// If exists, then delete the html file
	if _, err := os.Stat(path.Join(".", "cache", htmlFile)); err == nil {
		err = os.Remove(path.Join(".", "cache", mdFile))
		HandleErr(w, 500, err)
	}

}
