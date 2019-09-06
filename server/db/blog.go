package db

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// enums for doctype
var (
	DocTypeMD   int8 = 0
	DocTypeHTML int8 = 1
)

// Blog - Blog model type
type Blog struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	IDStr           string        `bson:"id_str" json:"id_str"`
	Title           string        `bson:"title" json:"title"`
	Description     string        `bson:"description" json:"description"`
	DescriptionLink string        `bson:"description_link" json:"description_link"`
	Author          string        `bson:"author" json:"author"`
	FormattedDate   string        `bson:"formatted_date" json:"formatted_date"`
	HTML            string        `bson:"html" json:"html"`
	Markdown        string        `bson:"markdown" json:"markdown"`
	DocType         int8          `bson:"doc_type" json:"doc_type"`
	Thumbnail       int8          `bson:"thumbnail" json:"thumbnail"`
	CreatedAt       int32         `bson:"created_at" json:"created_at"`
	IsTechnical     bool          `bson:"is_technical" json:"is_technical"`
	IsPublic        bool          `bson:"is_public" json:"is_public"`
	IsDeleted       bool          `bson:"is_deleted" json:"is_deleted"`
	IsSeries        bool          `bson:"is_series" json:"is_series"`
	SubBlogs        []SubBlog     `bson:"sub_blogs" json:"sub_blogs"`
}

// SubBlog - Blog model type
type SubBlog struct {
	Title         string `bson:"title" json:"title"`
	Description   string `bson:"description" json:"description"`
	FormattedDate string `bson:"formatted_date" json:"formatted_date"`
	HTML          string `bson:"html" json:"html"`
	Markdown      string `bson:"markdown" json:"markdown"`
	DocType       int8   `bson:"doc_type" json:"doc_type"`
}

// Blogs - Slice of Blogs
type Blogs []Blog

/*
Count returns the number of documents
that exists in the provided query
*/
func (bs *Blogs) Count(f bson.M) (int, error) {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return 0, err
	}
	defer ms.Close()

	return ms.DB(DBName).C("blogs").Find(f).Count()
}

/*
ReadAll reads all blog-documents from the database
and puts the values into `Blogs` slice
*/
func (bs *Blogs) ReadAll(f, s bson.M) error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("blogs").Find(f).Sort("-created_at").Select(s).All(bs)
}

/*
ReadFew reads a certain number of blog-documents from the database
the `skp` and `lim` values defines teh `skip` and `limit` values
for the query
*/
func (bs *Blogs) ReadFew(f, s bson.M, skp, lim int) error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("blogs").Find(f).Sort("-created_at").Select(s).Skip(skp).Limit(lim).All(bs)
}

// Create - Creates new Document
func (b *Blog) Create() error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	b.ID = bson.NewObjectId()
	return ms.DB(DBName).C("blogs").Insert(&b)
}

// Read - Reads single Document
func (b *Blog) Read(f, s bson.M) error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("blogs").Find(f).Select(s).One(b)
}

// Update - Updates a Document by ID
func (b *Blog) Update(s bson.M, u bson.M) error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err = ms.DB(DBName).C("blogs").Find(s).Apply(change, b)

	return err
}

// Delete - Deletes a Document
func (b *Blog) Delete(id bson.ObjectId) error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	// err := blogCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err = ms.DB(DBName).C("blogs").Find(bson.M{"_id": b.ID}).Apply(change, b)

	return err
}

// DeletePermanent - Deletes a document permanently
func (b *Blog) DeletePermanent() error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("blogs").Remove(bson.M{"_id": b.ID})
}
