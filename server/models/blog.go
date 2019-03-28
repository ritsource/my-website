package models

import (
	"fmt"
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
func (bs Blogs) Read(s bson.M) (Blogs, error) {
	fmt.Printf("bs %+v\n", bs)

	err := col.Find(s).All(&bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// Create - ..
func (b Blog) Create() (Blog, error) {
	err := col.Insert(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

// ReadSingle - ..
func (b Blog) ReadSingle(s bson.M) (Blog, error) {
	err := col.Find(s).One(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

// Update - ..
func (b Blog) Update(s bson.M, u bson.M) (Blog, error) {
	change := mgo.Change{
		Update: bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := col.Find(s).Apply(change, &b)
	
	return b, err
}

// Delete - ..
func (b Blog) Delete(id bson.ObjectId) (Blog, error) {
	// err := col.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := col.Find(bson.M{"_id": id}).Apply(change, &b)
	
	return b, err
}