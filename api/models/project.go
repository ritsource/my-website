package models

import (
	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Project - Project model type
type Project struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title           string        `bson:"title" json:"title"`
	Description     string        `bson:"description" json:"description"`
	DescriptionLink string        `bson:"description_link" json:"description_link"`
	HTML            string        `bson:"html" json:"html"`
	Markdown        string        `bson:"markdown" json:"markdown"`
	DocType         string        `bson:"doc_type" json:"doc_type"`
	Thumbnail       string        `bson:"thumbnail" json:"thumbnail"`
	Link            string        `bson:"link" json:"link"`
	CreatedAt       int32         `bson:"created_at" json:"created_at"`
	IsPublic        bool          `bson:"is_public" json:"is_public"`
	IsDeleted       bool          `bson:"is_deleted" json:"is_deleted"`
}

// MongoDB projectCollection for Files
var projectCol *mgo.Collection

func init() {
	projectCol = db.Client.DB(config.Secrets.DBName).C("projects")
}

// Projects - Slice of Projects
type Projects []Project

// Read - Reads all Documents
func (ps Projects) Read(s bson.M) (Projects, error) {
	err := projectCol.Find(s).Sort("-created_at").All(&ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// Create - Creates a new Document
func (p Project) Create() (Project, error) {
	err := projectCol.Insert(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

// Read - Reads single Document
func (p Project) Read(s bson.M) (Project, error) {
	err := projectCol.Find(s).One(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

// Update - Updates a Document
func (p Project) Update(s bson.M, u bson.M) (Project, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := projectCol.Find(s).Apply(change, &p)

	return p, err
}

// Delete - Deletes a document
func (p Project) Delete(id bson.ObjectId) (Project, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := projectCol.Find(bson.M{"_id": id}).Apply(change, &p)

	return p, err
}

// DeletePermanent - Deletes a document permanently
func (p Project) DeletePermanent(id bson.ObjectId) error {
	err := projectCol.Remove(bson.M{"_id": id})
	return err
}
