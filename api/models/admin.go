package models

import (
	"flag"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
)

// Admin (user) model type
type Admin struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string        `bson:"email" json:"email"`
	GoogleID string        `bson:"google_id" json:"id"`
}

// ACol - MongoDB Collection for Admin
var ACol *mgo.Collection

// BCol - MongoDB Collection for Blogs
var BCol *mgo.Collection

// PCol - MongoDB Collection for Projects
var PCol *mgo.Collection

func init() {
	var mdb *mgo.Database
	if flag.Lookup("test.v") == nil {
		mdb = db.Client.DB(config.Secrets.DBName) // mgo.DB for
	} else {
		mdb = db.Client.DB(config.Secrets.TestDBName) // mgo.DB for
	}

	ACol = mdb.C("admins")
	BCol = mdb.C("blogs")
	PCol = mdb.C("projects")
}

// Read - reads admin from Database (Not sure when I'm gonna need this)
func (a *Admin) Read(s bson.M) error {
	err := ACol.Find(s).One(a)
	if err != nil {
		return err
	}

	return nil
}

// Create - inserts a data to Admin collection on the Database
func (a *Admin) Create() error {
	err := ACol.Insert(a)
	if err != nil {
		return err
	}

	return nil
}
