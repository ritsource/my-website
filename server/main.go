package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/middleware"
	"github.com/ritwik310/my-website/server/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")

	r.HandleFunc("/auth/google", auth.GoogleLoginHandeler).Methods("GET")
	r.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/auth/current_user", middleware.CheckAuth(auth.CurrentUserHandler)).Methods("GET")

	r.HandleFunc("/admin/blog/all", middleware.CheckAuth(routes.ReadAllBlogs)).Methods("GET")
	r.HandleFunc("/admin/blog/{id}", middleware.CheckAuth(routes.ReadSingleBlog)).Methods("GET")
	r.HandleFunc("/admin/blog/new", middleware.CheckAuth(routes.CreateBlog)).Methods("POST")
	// r.HandleFunc("/admin/blog/edit/{id}", middleware.CheckAuth(routes.EditBlog)).Methods("PUT")
	r.HandleFunc("/admin/blog/edit/{id}", routes.EditBlog).Methods("PUT")
	r.HandleFunc("/admin/blog/delete/{id}", middleware.CheckAuth(routes.ReadSingleBlog)).Methods("DELETE")

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
