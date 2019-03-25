package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/mongo"
)

// "github.com/julienschmidt/httprouter"
// "go.mongodb.org/mongo-driver/mongo"

func main() {
	var err error

	// Connecting to Database (MongoDB)
	var mongoSession *mongo.Session // mongoSession ...
	mongoSession, err = mongo.NewSession(auth.Secrets.MongoURI)
	if err != nil {
		fmt.Printf("Error: unable to connect to mongo: %s\n", err)
	}

	defer mongoSession.Close()

	http.HandleFunc("/auth/current_user", auth.GetCurrentUserHandeler(mongoSession))
	http.HandleFunc("/auth/google", auth.HandleGoogleLogin)
	http.HandleFunc("/auth/google/callback", auth.GetGoogleCallbackHandeler(mongoSession))

	log.Fatal(http.ListenAndServe(":8080", nil))
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
