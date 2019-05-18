package models

import (
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
	IsMajor         bool          `bson:"is_major" json:"is_major"`
	IsPublic        bool          `bson:"is_public" json:"is_public"`
	IsDeleted       bool          `bson:"is_deleted" json:"is_deleted"`
}

// Projects - Slice of Projects
type Projects []Project

// Read - Reads all Documents
func (ps *Projects) Read(f, s bson.M) error {
	err := PCol.Find(s).Sort("-created_at").Select(s).All(ps)
	if err != nil {
		return err
	}

	return nil
}

// Create - Creates a new Document
func (p *Project) Create() error {
	err := PCol.Insert(p)
	if err != nil {
		return err
	}

	return nil
}

// Read - Reads single Document
func (p *Project) Read(f, s bson.M) error {
	err := PCol.Find(s).Select(s).One(p)
	if err != nil {
		return err
	}

	return nil
}

// Update - Updates a Document
func (p *Project) Update(s bson.M, u bson.M) error {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := PCol.Find(s).Apply(change, p)

	return err
}

// Delete - Deletes a document
func (p *Project) Delete(id bson.ObjectId) error {
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := PCol.Find(bson.M{"_id": p.ID}).Apply(change, p)

	return err
}

// DeletePermanent - Deletes a document permanently
func (p *Project) DeletePermanent(id bson.ObjectId) error {
	err := PCol.Remove(bson.M{"_id": p.ID})
	return err
}
