package main

import (
	"fmt"
	// "os"
	"log"
	"net/http"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/config"
)

var isDev bool
var client *mongo.Client

func init() {
	// Connecting to MongoDB
	var err error
	client, err = db.Connect(config.Secrets.MongoURI)
	if err != nil {
		fmt.Println("Error: MongoDB Connection")
		log.Fatal(err)
	}
}

func main() {
	// var err error

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"hello\": \"world\"}"))
	})

	mux.HandleFunc("/auth/current_user", auth.PickHandeler("/auth/current_user", client))
	mux.HandleFunc("/auth/google", auth.PickHandeler("/auth/google", client))
	mux.HandleFunc("/auth/google/callback", auth.PickHandeler("/auth/google/callback", client))

	// log.Fatal(http.ListenAndServe(":8080", nil))
	handler := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:3000"},
    AllowCredentials: true,
	}).Handler(mux)
	
	http.ListenAndServe(":8080", handler)
}