package models

import (
	// "fmt"
	// "encoding/json"
	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2"

	// "github.com/ritwik310/my-website/server/db"
	// "github.com/ritwik310/my-website/server/config"
)

// Project - Project model type
type Project struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	HTML string `bson:"html" json:"html"`
	Markdown string `bson:"markdown" json:"markdown"`
	Link string `bson:"link" json:"link"`
	ImageURL string `bson:"image_url" json:"image_url"`
	IsPublic string `bson:"is_public" json:"is_public"`
	IsDeleted string `bson:"is_deleted" json:"is_deleted"`
}