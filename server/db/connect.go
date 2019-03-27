package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ritwik310/my-website/server/config"
)

// Client - MongoDB connected Client
var Client *mongo.Client
var isDev bool

func init() {
	// Connecting to MongoDB
	var err error
	Client, err = connect(config.Secrets.MongoURI)

	if err != nil {
		fmt.Println("Error: couldn't connect to MongoDB")
		log.Fatal(err)
	}
}

// Connect - connects to MongoDB
func connect(uri string) (*mongo.Client, error) {
	var client *mongo.Client
	var err error

	// Connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return client, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return client, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}