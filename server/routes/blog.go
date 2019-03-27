package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	gContext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/models"
)

// CreateBlog ...
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	_, err := auth.CheckAuth(r, db.Client)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error:" + err.Error()))
		fmt.Println(err)
		return
	}

	var body models.Blog // to save Blog JSON body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// MongoDB Collection
	collection := db.Client.Database(config.Secrets.DatabaseName).Collection("blogs")

	// var newBlog models.Blog
	var result *mongo.InsertOneResult
	err = nil
	result, err = collection.InsertOne(context.TODO(), body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"unable to insert new admin\"}"))
		return
	}

	// Query the created Blog
	var newBlog models.Blog
	err = nil
	err = collection.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: result.InsertedID}}).Decode(&newBlog)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"created, but couldn't retreve data\"}"))
		return
	}

	// Marshaling result
	bData, err := newBlog.ToJSON()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"couldn't recover data\"}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadAllBlogs - read all blogs, both Public and Private
func ReadAllBlogs(w http.ResponseWriter, r *http.Request) {
	_ = gContext.Get(r, "admin")

	// MongoDB Collection
	collection := db.Client.Database(config.Secrets.DatabaseName).Collection("blogs")

	var allBlogs models.Blogs

	// Pass these options to the Find method
	findOptions := options.Find()
	// findOptions.SetLimit(2)

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem models.Blog
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		allBlogs = append(allBlogs, *&elem)
	}

	if err := cur.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// Marshaling result
	bData, err := allBlogs.ToJSON()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"couldn't recover data\"}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}
