package handlers

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ritwik310/my-website/server/models"
)

// Blog - ...
type Blog struct{
	Client *mongo.Client
	Db string
	Col string
}

// CreateOne ...
func (b Blog) CreateOne(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var blogObj models.Blog

	err := decoder.Decode(&blogObj)
	if err != nil {
			panic(err)
	}

	fmt.Printf("Blog: %+v\n", blogObj)
	collection := b.Client.Database(b.Db).Collection(b.Col)

	// var newBlog models.Blog
	var result *mongo.InsertOneResult
	err = nil
	result, err = collection.InsertOne(context.TODO(), blogObj)
	if err != nil {
		fmt.Printf("Error: unable to insert new admin %s\n", err)
		// return admin, err
	}

	fmt.Printf(":)))), %+v\n", result)
}