package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	// "reflect"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/models"
)

// MongoDB Collection for Blogs
var collection *mongo.Collection

func init() {
	collection = db.Client.Database(config.Secrets.DatabaseName).Collection("blogs")
}

// Writes Admin Un-Authenticated on Response
func writeError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\"}"))
	fmt.Println("Error:", err.Error())
}

// CreateBlog ...
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blogBody models.Blog // to save Blog JSON body
	var err error

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&blogBody)
	if err != nil {
		writeError(w, err, "Unable to read request body")
		return
	}

	// var newBlog models.Blog
	var result *mongo.InsertOneResult
	err = nil
	result, err = collection.InsertOne(context.TODO(), blogBody)
	if err != nil {
		writeError(w, err, "Unable to insert data")
		return
	}

	// Query the created Blog
	var newBlog models.Blog
	err = nil
	err = collection.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: result.InsertedID}}).Decode(&newBlog)
	if err != nil {
		writeError(w, err, "Insert successful, but unable to read")
		return
	}

	// Marshaling result
	bData, err := newBlog.ToJSON()
	if err != nil {
		writeError(w, err, "Insert successful, but unable to read")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadSingleBlog - ...
func ReadSingleBlog(w http.ResponseWriter, r *http.Request) {
	blogIDStr := mux.Vars(r)["id"] // Blog ObjectId String

	var theBlog models.Blog
	var err error

	var blogID primitive.ObjectID // Blog ObjectId (type ObjectId)
	blogID, err = primitive.ObjectIDFromHex(blogIDStr)

	// Query Blog
	err = collection.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: blogID}}).Decode(&theBlog)
	if err != nil {
		writeError(w, err, "Unable to query data")
		return
	}

	// Marshaling result
	bData, err := theBlog.ToJSON()
	if err != nil {
		writeError(w, err, "Unable to query data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// ReadAllBlogs - read all blogs, both Public and Private
func ReadAllBlogs(w http.ResponseWriter, r *http.Request) {
	var allBlogs models.Blogs

	// Pass these options to the Find method
	findOptions := options.Find()

	// Passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		writeError(w, err, "Unable to query data")
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
		if err != nil {
			writeError(w, err, "Unable to query data")
			return
		}
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// Marshaling result
	bData, err := allBlogs.ToJSON()
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

	// Takes ID string from URL Param and Turns it into MongoDB ObjectID
	// idStr := mux.Vars(r)["id"]
	bID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"]) // MongoDB ObjectID for the Blog
	if err != nil {
		writeError(w, err, "Unable to read request id")
		return
	}

	// Decoding request body from r.Body
	var body models.Blog

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		writeError(w, err, "Unable to read request body")
		return
	}

	// change := bson.D{
	// 	bson.E{Key: "_id", Value: blogID},
	// 	bson.E{Key: "_id", Value: blogID},
	// 	bson.E{Key: "_id", Value: blogID},
	// 	bson.E{Key: "_id", Value: blogID},
	// }

	// Query Blog
	err = nil
	result := collection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			bson.E{Key: "_id", Value: bID},
		},
		bson.D{
			bson.E{Key: "$set",
				Value: bson.E{
					bson.E{Key: "title", Value: body.Title},
					bson.E{Key: "description", Value: body.Description},
					bson.E{Key: "html", Value: body.HTML},
					bson.E{Key: "markdown", Value: body.Markdown},
					bson.E{Key: "image_url", Value: body.ImageURL},
				},
			},
		},
	)

	var doc models.Blog

	err = result.Decode(&doc)
	if err != nil {
		writeError(w, err, "Unable to read request body")
		return
	}

	// if err != nil {
	// 	writeError(w, err, "Unable to query data")
	// 	return
	// }

	// // Marshaling result
	// bData, err := theBlog.ToJSON()
	// if err != nil {
	// 	writeError(w, err, "Unable to query data")
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(SingleResult)

	fmt.Printf("Doc %+v\n", doc)
	w.Write([]byte("HELLO"))

}
