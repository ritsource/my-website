package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// Public Routes - Public routes are publically accessable (as the name suggests)
// it doesn't require admin authentication

// PubReadBlogs - Reads All Public Blog Data
func PubReadBlogs(w http.ResponseWriter, r *http.Request) {
	var bls models.Blogs
	var err error

	qTech := r.URL.Query()["tech"]

	if len(qTech) > 0 && qTech[0] == "true" {
		err = bls.Read(bson.M{"is_technical": true, "is_deleted": false, "is_public": true}, bson.M{}) // Read Public Blogs

	} else if len(qTech) > 0 && qTech[0] == "false" {
		err = bls.Read(bson.M{"is_technical": false, "is_deleted": false, "is_public": true}, bson.M{}) // Read Public Blogs

	} else {
		err = bls.Read(bson.M{"is_deleted": false, "is_public": true}, bson.M{}) // Read Public Blogs

	}

	HandleErr(w, 422, err)

	WriteData(w, bls) // Write Data
}

// PubReadBlog - Read Public blog by ID
func PubReadBlog(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog
	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Read blog
	err := bl.Read(bson.M{"_id": bson.ObjectIdHex(bIDStr), "is_deleted": false, "is_public": true}, bson.M{})
	HandleErr(w, 422, err)

	WriteData(w, bl) // Write Data
}

// PubReadProjects -
func PubReadProjects(w http.ResponseWriter, r *http.Request) {
	var pr models.Projects

	err := pr.Read(bson.M{"is_deleted": false, "is_public": true}, bson.M{}) // Read from DB
	HandleErr(w, 422, err)

	WriteData(w, pr) // Write Data
}

// PubReadProject - Read Public Project By ID
func PubReadProject(w http.ResponseWriter, r *http.Request) {
	var pr models.Project
	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Read project
	err := pr.Read(bson.M{"_id": bson.ObjectIdHex(pIDStr), "is_deleted": false, "is_public": true}, bson.M{})
	HandleErr(w, 422, err)

	WriteData(w, pr) // Write Data
}
