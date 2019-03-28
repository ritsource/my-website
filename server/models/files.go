package models

import (
	// "fmt"
	// "encoding/json"
	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2"

	// "github.com/ritwik310/my-website/server/db"
	// "github.com/ritwik310/my-website/server/config"
)

// Files - Files model type
type Files struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Type string `bson:"type" json:"type"`
	FileURL string `bson:"file_url" json:"file_url"`
	IsPublic string `bson:"is_public" json:"is_public"`
	IsDeleted string `bson:"is_deleted" json:"is_deleted"`
}