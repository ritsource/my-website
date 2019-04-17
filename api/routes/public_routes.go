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

	bls, err := bls.Read(bson.M{"is_deleted": false, "is_public": true}) // Read Public Blogs
	HandleErr(w, 422, err)

	WriteData(w, bls) // Write Data
}

// PubReadBlog - Read Public blog by ID
func PubReadBlog(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog
	bIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	// Read blog
	bl, err := bl.ReadSingle(bson.M{"_id": bson.ObjectIdHex(bIDStr), "is_deleted": false, "is_public": true})
	HandleErr(w, 422, err)

	WriteData(w, bl) // Write Data
}

// PubReadProjects -
func PubReadProjects(w http.ResponseWriter, r *http.Request) {
	var pr models.Projects

	pr, err := pr.Read(bson.M{"is_deleted": false, "is_public": true}) // Read from DB
	HandleErr(w, 422, err)

	WriteData(w, pr) // Write Data
}

// PubReadProject - Read Public Project By ID
func PubReadProject(w http.ResponseWriter, r *http.Request) {
	var pr models.Project
	pIDStr := mux.Vars(r)["id"] // Project ObjectId String

	// Read project
	pr, err := pr.ReadSingle(bson.M{"_id": bson.ObjectIdHex(pIDStr), "is_deleted": false, "is_public": true})
	HandleErr(w, 422, err)

	WriteData(w, pr) // Write Data
}
