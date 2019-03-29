package models

// Files aren't needed for NOW.........

import (
	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// File - Files model type
type File struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Type        string        `bson:"type" json:"type"`
	Extension 	string 				`bson:"extension" josn:"extension"`
	FileURL     string        `bson:"file_url" json:"file_url"`
	IsPublic    bool          `bson:"is_public" json:"is_public"`
	IsDeleted   bool          `bson:"is_deleted" json:"is_deleted"`
}

// MongoDB fileCollection for Files
var fileCol *mgo.Collection

func init() {
	fileCol = db.Client.DB(config.Secrets.DBName).C("files")
}

// Files - Slice of Files
type Files []File

// ReadAll - ..
func (fs Files) Read(s bson.M) (Files, error) {
	err := fileCol.Find(s).All(&fs)
	if err != nil {
		return nil, err
	}

	return fs, nil
}

// Create - ..
func (f File) Create() (File, error) {
	err := fileCol.Insert(&f)
	if err != nil {
		return f, err
	}

	return f, nil
}

// ReadSingle - ..
func (f File) ReadSingle(s bson.M) (File, error) {
	err := fileCol.Find(s).One(&f)
	if err != nil {
		return f, err
	}

	return f, nil
}

// Update - ..
func (f File) Update(s bson.M, u bson.M) (File, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := fileCol.Find(s).Apply(change, &f)

	return f, err
}

// Delete - ..
func (f File) Delete(id bson.ObjectId) (File, error) {
	// err := fileCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := fileCol.Find(bson.M{"_id": id}).Apply(change, &f)

	return f, err
}
