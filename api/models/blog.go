package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
)

// Blog - Blog model type
type Blog struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title         string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	Author        string        `bson:"author" json:"author"`
	FormattedDate string        `bson:"formatted_date" json:"formatted_date"`
	HTML          string        `bson:"html" json:"html"`
	Markdown      string        `bson:"markdown" json:"markdown"`
	DocType       string        `bson:"doc_type" json:"doc_type"`
	IsPublic      bool          `bson:"is_public" json:"is_public"`
	IsDeleted     bool          `bson:"is_deleted" json:"is_deleted"`
}

// MongoDB blogCollection for Blogs
var blogCol *mgo.Collection

func init() {
	blogCol = db.Client.DB(config.Secrets.DBName).C("blogs")
}

// Blogs - Slice of Blogs
type Blogs []Blog

// ReadAll - ..
func (bs Blogs) Read(s bson.M) (Blogs, error) {
	err := blogCol.Find(s).All(&bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// Create - ..
func (b Blog) Create() (Blog, error) {
	err := blogCol.Insert(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

// ReadSingle - ..
func (b Blog) ReadSingle(s bson.M) (Blog, error) {
	err := blogCol.Find(s).One(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

// Update - ..
func (b Blog) Update(s bson.M, u bson.M) (Blog, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := blogCol.Find(s).Apply(change, &b)

	return b, err
}

// Delete - ..
func (b Blog) Delete(id bson.ObjectId) (Blog, error) {
	// err := blogCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := blogCol.Find(bson.M{"_id": id}).Apply(change, &b)

	return b, err
}

// DeletePermanent - Deletes a document permanently
func (b Blog) DeletePermanent(id bson.ObjectId) error {
	err := blogCol.Remove(bson.M{"_id": id})
	return err
}
