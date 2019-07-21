package db

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

// MongoConn session, to communicate with mongo server
var MongoConn *mgo.Session

// MgoDB can be used to manipulate data in the database
var MgoDB *mgo.Database

func init() {
	err := Connect()
	if err != nil {
		logrus.Panicf("%v\n", err)
	}

	//
	if flag.Lookup("test.v") == nil {
		MgoDB = MongoConn.DB(os.Getenv("DB_NAME"))
	} else {
		MgoDB = MongoConn.DB(os.Getenv("TEST_DB_NAME"))
	}

}

// Connect creates a connection to mongoDB server
func Connect() error {
	var err error

	// dialing to mongoURI
	MongoConn, err = mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		logrus.Errorf("couldn't connect to mongodb server\n")
		return err
	}

	logrus.Infof("connected to MongoDB!")
	return nil
}

// Close closes connection to mongoDB server
func Close() {
	MongoConn.Close()
}
