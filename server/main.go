package main

import (
	"fmt"
	// "log"
	"net/http"

	"github.com/rs/cors"

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


	mux := http.NewServeMux()

	mux.HandleFunc("/auth/current_user", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"hello\": \"world\"}"))
	})
	

	// mux.HandleFunc("/auth/current_user", auth.GetCurrentUserHandeler(mongoSession))
	mux.HandleFunc("/auth/google", auth.HandleGoogleLogin)
	mux.HandleFunc("/auth/google/callback", auth.GetGoogleCallbackHandeler(mongoSession))

	// log.Fatal(http.ListenAndServe(":8080", nil))
	handler := cors.Default().Handler(mux)
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
