package routes

import (
	"encoding/json"
	"io"
	// "io/ioutil"
	"fmt"
	"os"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// DownloadFile -
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

// CreateProject ...
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var body models.Project // to save Project JSON body
	var err error

	// Decoding request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		WriteError(w, 422, err, "Unable to read request body")
		return
	}

	// Inserting Document
	nProject, err := body.Create()
	if err != nil {
		WriteError(w, 422, err, "Failed to insert new document")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nProject)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// GetProjectDocument - 
// This Handler Queries Project from teh database first, then checks
// if file exists or not in the "/satic" folder, (K8s PV in Prod),
// If file exists then serves it else redirects the user to the source
// saved in the mongo document, also if file doesn't exist the it downloads
// the file in the static folder
func GetProjectDocument(w http.ResponseWriter, r *http.Request) {
	var err error
	var mProject models.Project

	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Read project
	mProject, err = mProject.ReadSingle(bson.M{"_id": bson.ObjectIdHex(pIDStr)})
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Generating Local file Name
	var fileName string
	if mProject.DocType == "markdown" {
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
	if mProject.DocType == "markdown" {
		srcFilePath = mProject.Markdown
	} else {
		srcFilePath = mProject.HTML
	}

	// Redirecting to the source file
	http.Redirect(w, r, srcFilePath, http.StatusTemporaryRedirect)
	
	// Downloading the file
	DownloadFile("./static/" + fileName, srcFilePath)
}

// ReadOneProject - ...
func ReadOneProject(w http.ResponseWriter, r *http.Request) {
	var err error
	var mProject models.Project

	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Read project
	mProject, err = mProject.ReadSingle(bson.M{"_id": bson.ObjectIdHex(pIDStr)})
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

// ReadProjects - read all projects, both Public and Private
func ReadProjects(w http.ResponseWriter, r *http.Request) {
	var err error
	var mProjects models.Projects

	// Read Project
	mProjects, err = mProjects.Read(bson.M{})
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

// EditProject - ...
func EditProject(w http.ResponseWriter, r *http.Request) {
	var err error
	var nProject models.Project

	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Decoding request body
	var body map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		WriteError(w, 422, err, "Unable to read request body")
		return
	}

	fmt.Printf("body %+v\n", body)

	// Update Project Document
	nProject, err = nProject.Update(
		bson.M{"_id": bson.ObjectIdHex(pIDStr)},
		body,
	)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nProject)
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// DeleteProject - ...
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var err error
	var dProject models.Project

	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Read Project
	_, err = dProject.Delete(bson.ObjectIdHex(pIDStr))
	if err != nil {
		WriteError(w, 422, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
