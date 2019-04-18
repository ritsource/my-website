package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
)

// Admin - admin (user) model type
type Admin struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string        `bson:"email" json:"email"`
	GoogleID string        `bson:"google_id" json:"id"`
}

// MongoDB Collection for Blogs
var adminCol *mgo.Collection

func init() {
	adminCol = db.Client.DB(config.Secrets.DBName).C("admins")
}

// ReadAll - ..
func (a Admin) Read(s bson.M) (Admin, error) {
	err := adminCol.Find(s).One(&a)
	if err != nil {
		return a, err
	}

	return a, nil
}

// Create - ..
func (a Admin) Create() (Admin, error) {
	err := adminCol.Insert(&a)
	if err != nil {
		return a, err
	}

	return a, nil
}
