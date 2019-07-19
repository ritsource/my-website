package db

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
	IsSeries        bool          `bson:"is_series" json:"is_series"`
	SubBlogs        []SubBlog     `bson:"sub_blogs" json:"sub_blogs"`
}

// SubBlog - Blog model type
type SubBlog struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title         string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	FormattedDate string        `bson:"formatted_date" json:"formatted_date"`
	HTML          string        `bson:"html" json:"html"`
	Markdown      string        `bson:"markdown" json:"markdown"`
	DocType       string        `bson:"doc_type" json:"doc_type"`
}

// Blogs - Slice of Blogs
type Blogs []Blog

// Read - Reads all Documents from blogs
func (bs *Blogs) Read(f, s bson.M) error {
	err := MgoDB.C("blogs").Find(f).Sort("-created_at").Select(s).All(bs)
	if err != nil {
		return err
	}

	return nil
}

// Create - Creates new Document
func (b *Blog) Create() error {
	err := MgoDB.C("blogs").Insert(&b)
	if err != nil {
		return err
	}

	return nil
}

// Read - Reads single Document
func (b *Blog) Read(f, s bson.M) error {
	err := MgoDB.C("blogs").Find(f).Select(s).One(b)
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
	_, err := MgoDB.C("blogs").Find(s).Apply(change, b)

	return err
}

// Delete - Deletes a Document
func (b *Blog) Delete(id bson.ObjectId) error {
	// err := blogCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := MgoDB.C("blogs").Find(bson.M{"_id": b.ID}).Apply(change, b)

	return err
}

// DeletePermanent - Deletes a document permanently
func (b *Blog) DeletePermanent() error {
	err := MgoDB.C("blogs").Remove(bson.M{"_id": b.ID})
	return err
}
