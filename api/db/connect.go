package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	"github.com/ritwik310/my-website/api/config"
)

// Client - MongoDB connected Client
var Client *mgo.Session

func init() {
	// Connecting to MongoDB
	var err error
	Client, err = mgo.Dial(config.Secrets.MongoURI)

	if err != nil {
		fmt.Println("Error: couldn't connect to MongoDB")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// Close - closes, the db
func Close() {
	Client.Close()
}
