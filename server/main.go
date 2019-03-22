package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ritwik310/my-website/server/auth"
)

// "github.com/julienschmidt/httprouter"
// "go.mongodb.org/mongo-driver/mongo"

func main() {
	http.HandleFunc("/auth/google", auth.HandleGoogleLogin)
	http.HandleFunc("/auth/google/callback", auth.HandleGoogleCallback)

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
