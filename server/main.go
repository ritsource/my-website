package main

import (
	"fmt"
	// "os"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/handlers"
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
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"hello\": \"world\"}"))
	})

	router.GET("/auth/current_user", auth.PickHandeler("/auth/current_user", client))
	router.GET("/auth/google", auth.PickHandeler("/auth/google", client))
	router.GET("/auth/google/callback", auth.PickHandeler("/auth/google/callback", client))

	blog := handlers.Blog{
		Client: client,
		Db: "dev_db",
		Col: "blogs",
	}

	// router.READ("/admin/blog/:id", )
	router.POST("/admin/add_blog", blog.CreateOne)
	// router.PUT("/admin/edit_blog/:id", )
	// router.DELETE("/admin/delete_blog/:id", )


	// log.Fatal(http.ListenAndServe(":8080", nil))
	handler := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:3000"},
    AllowCredentials: true,
	}).Handler(router)
	
	http.ListenAndServe(":8080", handler)
}