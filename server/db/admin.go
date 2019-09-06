package db

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
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
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("admins").Find(s).One(a)
}

// Create creates a new admin document
func (a *Admin) Create() error {
	ms, err := mgo.Dial(MongoURI)
	if err != nil {
		logrus.Printf("Could not connect to mongo: %v\n", err)
		return err
	}
	defer ms.Close()

	return ms.DB(DBName).C("admins").Insert(a)
}
