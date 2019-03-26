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

// "github.com/julienschmidt/httprouter"

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

// Utils
// func Message(status bool, message string) map[string]interface{} {
// 	return map[string]interface{}{"status": status, "message": message}
// }

// func Respond(w http.ResponseWriter, data map[string]interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }

// Response Haldelers
func handler(w http.ResponseWriter, r *http.Request) {
	// tokenHeader := r.Header.Get("Authorization")
	fmt.Fprintf(w, "Hi there, Home here!")
}


	// // Mux Router
	// router := mux.NewRouter()
	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Access-Control-Allow-Origin"}) // Request Headers
	// methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}) // Allowed Methods
	// origins := handlers.AllowedOrigins([]string{"*", "http://localhost:3000"}) // Allowed Origins

	// // Route Handelers
	// router.HandleFunc("/auth/current_user", auth.GetCurrentUserHandeler(mongoSession)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/auth/google", auth.HandleGoogleLogin).Methods("GET", "OPTIONS")
	// router.HandleFunc("/auth/google/callback", auth.GetGoogleCallbackHandeler(mongoSession)).Methods("GET", "OPTIONS")

	// // Server
	// log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
