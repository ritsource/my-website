package db

import (
	"gopkg.in/mgo.v2/bson"
)

// Admin .
type Admin struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string        `bson:"email" json:"email"`
	GoogleID string        `bson:"google_id" json:"id"`
}

// Read reads from the dataabse
func (a *Admin) Read(s bson.M) error {
	err := MgoDB.C("admins").Find(s).One(a)
	if err != nil {
		return err
	}

	return nil
}

// Create creates a new admin document
func (a *Admin) Create() error {
	err := MgoDB.C("admins").Insert(a)
	if err != nil {
		return err
	}

	return nil
}
