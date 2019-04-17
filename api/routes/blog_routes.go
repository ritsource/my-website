package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

// Content -
type Content interface {
	Create() (Content, error)
	ReadSingle() (Content, error)
	Update() (Content, error)
	Delete() (Content, error)
	DeletePermanent() error
}

// WriteErr -
func WriteErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)                                     // Status Code
	w.Header().Set("Content-Type", "application/json")        // Response Type - JSON
	w.Write([]byte("{\"message\": \"" + err.Error() + "\"}")) // Error Message

	fmt.Println("Error:", err.Error())
}

// HandleErr -
func HandleErr(w http.ResponseWriter, status int, err error) {
	if err != nil {
		WriteErr(w, status, err)
		return
	}
}

// WriteData -
func WriteData(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	HandleErr(w, 500, err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// CreateBlog -
func (bl Content) CreateHandler(w http.ResponseWriter, r *http.Request) {
	// var bl models.Blog // Blog

	decoder := json.NewDecoder(r.Body) // Read JSON Body
	err := decoder.Decode(&bl)
	HandleErr(w, 500, err)

	bl.CreatedAt = int32(time.Now().Unix()) // Set Creation Time

	_, err = bl.Create() // Create Document in the Database
	HandleErr(w, 422, err)

	// Redirecting to All-Blogs route handler
	http.Redirect(w, r, "/admin/blog/all", 302) // 302 - POST to GET
}

// ReadBlog -
func ReadBlog(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog          // Blog
	bIDStr := mux.Vars(r)["id"] // Blog ObjectID (String)

	bl, err := bl.ReadSingle(bson.M{"_id": bson.ObjectIdHex(bIDStr)}) // Read Document
	HandleErr(w, 442, err)

	WriteData(w, bl) // Write the Data
}

// ReadBlogs -
func ReadBlogs(w http.ResponseWriter, r *http.Request) {
	var bls models.Blogs // Blogs or []Blog

	bls, err := bls.Read(bson.M{}) // Read all Blogs bson.M{}
	HandleErr(w, 442, err)

	WriteData(w, bls) // Write Data
}

// EditBlog -
func EditBlog(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog              // Blog
	var body map[string]interface{} // because cannot use bl (type models.Blog) as type bson.M in argument to bl.Update
	bIDStr := mux.Vars(r)["id"]     // Blog ObjectID (String)

	decoder := json.NewDecoder(r.Body) // Read Request JSON
	err := decoder.Decode(&body)
	HandleErr(w, 422, err)

	bl, err = bl.Update(bson.M{"_id": bson.ObjectIdHex(bIDStr)}, body) // Update Document in Database
	HandleErr(w, 500, err)

	WriteData(w, bl) // Write Data
}

// DeleteBlog - Deletes a Blog (Not Permanently)
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog          // Blog
	bIDStr := mux.Vars(r)["id"] // Blog ObjectID (String)

	bl, err := bl.Delete(bson.ObjectIdHex(bIDStr)) // Editing Document
	HandleErr(w, 422, err)

	WriteData(w, bl) // Writing Data
}

// DeleteBlogF - Deletes a blog Permanently
func DeleteBlogF(w http.ResponseWriter, r *http.Request) {
	var bl models.Blog          // Blog
	bIDStr := mux.Vars(r)["id"] // Blog ObjectID (String)

	err := bl.DeletePermanent(bson.ObjectIdHex(bIDStr)) // Deleting Document
	HandleErr(w, 422, err)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}
