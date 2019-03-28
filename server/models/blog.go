package models

import (
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/config"
)

// Blog - Blog model type
type Blog struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	HTML string `bson:"html" json:"html"`
	Markdown string `bson:"markdown" json:"markdown"`
	ImageURL string `bson:"image_url" json:"image_url"`
	IsPublic bool `bson:"is_public" json:"is_public"`
	IsDeleted bool `bson:"is_deleted" json:"is_deleted"`
}

// MongoDB Collection for Blogs
var col *mgo.Collection

func init() {
	col = db.Client.DB(config.Secrets.DatabaseName).C("blogs")
}

// Blogs - Slice of Blogs
type Blogs []Blog

// ReadAll - ..
func (bs Blogs) ReadAll(s bson.M) (*Blogs, error) {
	err := col.Find(s).All(&bs)
	if err != nil {
		return nil, err
	}

	return &bs, nil
}

// Create - ..
func (b Blog) Create() (*Blog, error) {
	err := col.Insert(&b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// ReadSingle - ..
func (b Blog) ReadSingle(s bson.M) (*Blog, error) {
	err := col.Find(s).One(&b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// Update - ..
func (b Blog) Update(s bson.M, u bson.M) (*Blog, error) {
	err := col.Update(s, bson.M{"$set": u})
	if err != nil {
		return nil, err
	}

	return b.ReadSingle(s)
}

// ToJSON - ...
func (b Blog) ToJSON() ([]byte, error) {
	var bData []byte
	var err error
	
	bData, err = json.Marshal(b)
	if err != nil {
		fmt.Println("Error: toJSON error:", err)
	}

	return bData, err
}