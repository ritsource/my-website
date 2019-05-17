package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// Blogs - Slice of Blogs
type Blogs []Blog

// Read - Reads all Documents from blogs
func (bs *Blogs) Read(f, s bson.M) error {
	err := BCol.Find(f).Sort("-created_at").Select(s).All(bs)
	if err != nil {
		return err
	}

	return nil
}

// Create - Creates new Document
func (b *Blog) Create() error {
	err := BCol.Insert(&b)
	if err != nil {
		return err
	}

	return nil
}

// Read - Reads single Document
func (b *Blog) Read(f, s bson.M) error {
	err := BCol.Find(f).Select(s).One(b)
	if err != nil {
		return err
	}

	return nil
}

// Update - Updates a Document by ID
func (b *Blog) Update(s bson.M, u bson.M) error {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := BCol.Find(s).Apply(change, b)

	return err
}

// Delete - Deletes a Document
func (b *Blog) Delete(id bson.ObjectId) error {
	// err := blogCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := BCol.Find(bson.M{"_id": b.ID}).Apply(change, b)

	return err
}

// DeletePermanent - Deletes a document permanently
func (b *Blog) DeletePermanent(id bson.ObjectId) error {
	err := BCol.Remove(bson.M{"_id": b.ID})
	return err
}
