package routes

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	// "reflect"

	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"


	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2"

	// "github.com/ritwik310/my-website/server/config"
	// "github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/models"
)

// // MongoDB Collection for Blogs
// var collection *mongo.Collection

// func init() {
// 	collection := db.Client.DB(config.Secrets.DatabaseName).C("blogs")
// }

// Writes Admin Un-Authenticated on Response
func writeError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\"}"))
	fmt.Println("Error:", err.Error())
}

// CreateBlog ...
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var body models.Blog // to save Blog JSON body
	var err error

	// Decoding request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		writeError(w, err, "Unable to read request body")
		return
	}

	// Inserting Document
	nBlog, err := body.Create()
	if err != nil {
		writeError(w, err, "Failed to insert new document")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nBlog)
	if err != nil {
		writeError(w, err, "Unable to query data")
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
		writeError(w, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mBlog)
	if err != nil {
		writeError(w, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadBlogs - read all blogs, both Public and Private
func ReadBlogs(w http.ResponseWriter, r *http.Request) {
	var err error
	var mBlogs models.Blogs

	// Read blog
	mBlogs, err = mBlogs.Read(bson.M{})
	if err != nil {
		writeError(w, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(mBlogs)
	if err != nil {
		writeError(w, err, "Unable to query data")
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
		writeError(w, err, "Unable to read request body")
		return
	}

	fmt.Printf("body %+v\n", body)

	// Update Blog Document
	nBlog, err = nBlog.Update(
		bson.M{"_id": bson.ObjectIdHex(bIDStr)},
		body,
	)	
	if err != nil {
		writeError(w, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := json.Marshal(nBlog)
	if err != nil {
		writeError(w, err, "Unable to query data")
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
		writeError(w, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Successfully deleted\"}"))
}