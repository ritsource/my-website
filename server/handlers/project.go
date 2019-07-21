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

// CreateProject creates a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// reading JSON body
	var p db.Project

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// creation time
	p.CreatedAt = int32(time.Now().Unix())

	// insert new document
	err = p.Create()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// redirecting to `projects` route handler
	http.Redirect(w, r, "/api/private/projects", http.StatusTemporaryRedirect) // 302 - POST to GET
}

// ReadProject reads a single project
func ReadProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no project-id provided"))
		return
	}

	// reading document
	var p db.Project
	err := p.Read(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{})
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
	writeJSON(w, p)
}

// ReadProjects reads all projects
func ReadProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	var ps db.Projects
	err := ps.Read(bson.M{}, bson.M{})
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	writeJSON(w, ps)
}

// EditProject dits a project by `_id`
func EditProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no project-id provided"))
		return
	}

	// reading request body
	var body map[string]interface{} // because cannot use b (type db.Project) as type bson.M in argument to bl.Update
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// editing document
	var p db.Project
	err = p.Update(bson.M{"_id": bson.ObjectIdHex(id)}, body) // Update Document in Database
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	writeJSON(w, p) // Write Data
}

// DeleteProject - Deletes a project (Not Permanently)
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no project-id provided"))
		return
	}

	// editing document to `is_deleted: true`
	var p db.Project
	err := p.Delete(bson.ObjectIdHex(id))
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	// writing the updated data
	writeJSON(w, p)
}

// DeleteProjectPrem - Deletes a project Permanently
func DeleteProjectPrem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no project-id provided"))
		return
	}

	// deleting document (permanently)
	var p db.Project
	p.ID = bson.ObjectIdHex(id)
	err := p.DeletePermanent()
	if err != nil {
		writeErr(w, 422, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
