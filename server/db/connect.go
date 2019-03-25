package db

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - connects to MongoDB
func Connect(uri string) (*mongo.Client, error) {
	var client *mongo.Client
	var err error

	// Connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return client, err
	}

	defer cancel()

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return client, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}