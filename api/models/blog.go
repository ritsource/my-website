package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
)

// Blog - Blog model type
type Blog struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title           string        `bson:"title" json:"title"`
	Description     string        `bson:"description" json:"description"`
	DescriptionLink string        `bson:"description_link" json:"description_link"`
	Author          string        `bson:"author" json:"author"`
	FormattedDate   string        `bson:"formatted_date" json:"formatted_date"`
	HTML            string        `bson:"html" json:"html"`
	Markdown        string        `bson:"markdown" json:"markdown"`
	DocType         string        `bson:"doc_type" json:"doc_type"`
	Thumbnail       string        `bson:"thumbnail" json:"thumbnail"`
	CreatedAt       int32         `bson:"created_at" json:"created_at"`
	IsTechnical     bool          `bson:"is_technical" json:"is_technical"`
	IsPublic        bool          `bson:"is_public" json:"is_public"`
	IsDeleted       bool          `bson:"is_deleted" json:"is_deleted"`
}

var blogCol *mgo.Collection // MongoDB blogCollection for Blogs

func init() {
	blogCol = db.Client.DB(config.Secrets.DBName).C("blogs")
}

// Blogs - Slice of Blogs
type Blogs []Blog

// Read - Reads all Documents from blogs
func (bs Blogs) Read(s bson.M) (Blogs, error) {
	err := blogCol.Find(s).Sort("-created_at").All(&bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// Create - Creates new Document
func (b Blog) Create() (Blog, error) {
	err := blogCol.Insert(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

// Read - Reads single Document
func (b Blog) Read(s bson.M) (Blog, error) {
	err := blogCol.Find(s).One(&b)
	if err != nil {
		return b, err
	}

	return b, nil
}

// Update - Updates a Document by ID
func (b Blog) Update(s bson.M, u bson.M) (Blog, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := blogCol.Find(s).Apply(change, &b)

	return b, err
}

// Delete - Deletes a Document
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
