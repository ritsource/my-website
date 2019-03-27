package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/config"
	midd "github.com/ritwik310/my-website/server/middleware"
	"github.com/ritwik310/my-website/server/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")

	r.HandleFunc("/auth/current_user", midd.AuthRequired(auth.CurrentUserHandler)).Methods("GET")
	r.HandleFunc("/auth/google", auth.GoogleLoginHandeler).Methods("GET")
	r.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler).Methods("GET")

	r.HandleFunc("/admin/blogs", midd.AuthRequired(routes.ReadAllBlogs)).Methods("GET")
	r.HandleFunc("/admin/blog/{id}", indexHandler).Methods("GET")
	r.HandleFunc("/admin/add_blog", midd.AuthRequired(routes.CreateBlog)).Methods("POST")
	r.HandleFunc("/admin/edit_blog", indexHandler).Methods("PUT")
	r.HandleFunc("/admin/delete_blog", indexHandler).Methods("DELETE")

	ch := cors.New(cors.Options{
		AllowedOrigins:   config.Secrets.AllowedCorsURLs,
		AllowCredentials: true,
	}).Handler(r)

	log.Fatal(http.ListenAndServe(":8080", ch))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"hello world\", \"from\": \"Ritwik :)\"}"))
}
