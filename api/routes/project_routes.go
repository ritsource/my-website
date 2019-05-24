package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateProject -
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var pr models.Project // Project

	decoder := json.NewDecoder(r.Body) // Read JSON Body
	err := decoder.Decode(&pr)
	HandleErr(w, 500, err)

	pr.CreatedAt = int32(time.Now().Unix()) // Set Creation Time

	err = pr.Create() // Create Document in the Database
	HandleErr(w, 422, err)

	// Redirecting to All-Projects route handler
	http.Redirect(w, r, "/api/private/project/all", 302) // 302 - POST to GET
}

// ReadProject -
func ReadProject(w http.ResponseWriter, r *http.Request) {
	var pr models.Project       // Project
	pIDStr := mux.Vars(r)["id"] // Project ObjectID (String)

	err := pr.Read(bson.M{"_id": bson.ObjectIdHex(pIDStr)}, bson.M{}) // Read Document
	HandleErr(w, 442, err)

	WriteData(w, pr) // Write the Data
}

// ReadProjects -
func ReadProjects(w http.ResponseWriter, r *http.Request) {
	var prs models.Projects // Projects or []Project

	err := prs.Read(bson.M{}, bson.M{}) // Read all Projects bson.M{}
	HandleErr(w, 442, err)

	WriteData(w, prs) // Write Data
}

// EditProject -
func EditProject(w http.ResponseWriter, r *http.Request) {
	var pr models.Project           // Project
	var body map[string]interface{} // because cannot use pr (type models.Project) as type bson.M in argument to pr.Update
	pIDStr := mux.Vars(r)["id"]     // Project ObjectID (String)

	decoder := json.NewDecoder(r.Body) // Read Request JSON
	err := decoder.Decode(&body)
	HandleErr(w, 422, err)

	err = pr.Update(bson.M{"_id": bson.ObjectIdHex(pIDStr)}, body) // Update Document in Database
	HandleErr(w, 500, err)

	WriteData(w, pr) // Write Data
}

// DeleteProject - Deletes a Project (Not Permanently)
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var pr models.Project       // Project
	pIDStr := mux.Vars(r)["id"] // Project ObjectID (String)

	err := pr.Delete(bson.ObjectIdHex(pIDStr)) // Editing Document
	HandleErr(w, 422, err)

	WriteData(w, pr) // Writing Data
}

// DeleteProjectF - Deletes a project Permanently
func DeleteProjectF(w http.ResponseWriter, r *http.Request) {
	var pr models.Project       // Project
	pIDStr := mux.Vars(r)["id"] // Project ObjectID (String)

	pr.ID = bson.ObjectIdHex(pIDStr)
	err := pr.DeletePermanent() // Deleting Document
	HandleErr(w, 422, err)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
