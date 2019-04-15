package models

import (
	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Project - Project model type
type Project struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	HTML        string        `bson:"html" json:"html"`
	Markdown    string        `bson:"markdown" json:"markdown"`
	DocType     string        `bson:"doc_type" json:"doc_type"`
	Link        string        `bson:"link" json:"link"`
	CreatedAt   int32         `bson:"created_at" json:"created_at"`
	IsPublic    bool          `bson:"is_public" json:"is_public"`
	IsDeleted   bool          `bson:"is_deleted" json:"is_deleted"`
}

// MongoDB projectCollection for Files
var projectCol *mgo.Collection

func init() {
	projectCol = db.Client.DB(config.Secrets.DBName).C("projects")
}

// Projects - Slice of Projects
type Projects []Project

// ReadAll - ..
func (ps Projects) Read(s bson.M) (Projects, error) {
	err := projectCol.Find(s).Sort("-created_at").All(&ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// Create - ..
func (p Project) Create() (Project, error) {
	err := projectCol.Insert(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

// ReadSingle - ..
func (p Project) ReadSingle(s bson.M) (Project, error) {
	err := projectCol.Find(s).One(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

// Update - ..
func (p Project) Update(s bson.M, u bson.M) (Project, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := projectCol.Find(s).Apply(change, &p)

	return p, err
}

// Delete - ..
func (p Project) Delete(id bson.ObjectId) (Project, error) {
	// err := projectCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
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
