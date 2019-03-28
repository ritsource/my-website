package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

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
