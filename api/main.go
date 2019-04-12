package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ritwik310/my-website/api/backup"
	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
	"github.com/ritwik310/my-website/api/middleware"
	"github.com/ritwik310/my-website/api/routes"
)

func main() {
	// Inserts Data to a Backup Database, once every day..
	go backup.BackItup()

	// Closes Database, on main.go File Exit
	defer db.Close()

	// Router (Gorilla Mux)
	r := mux.NewRouter()

	// Takes care of Static file serving
	var dir string
	flag.StringVar(&dir, "dir", "./cache", "usage")
	flag.Parse()

	// Static file server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// Dummy handler
	r.HandleFunc("/", indexHandler).Methods("GET")

	// Authentication (Admin) handlers
	r.HandleFunc("/auth/google", routes.GoogleLoginHandeler).Methods("GET")
	r.HandleFunc("/auth/google/callback", routes.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/auth/current_user", middleware.CheckAuth(routes.CurrentUserHandler)).Methods("GET")

	// Admin, Blog Handlers
	r.HandleFunc("/admin/blog/all", middleware.CheckAuth(routes.ReadBlogs)).Methods("GET")
	r.HandleFunc("/admin/blog/{id}", middleware.CheckAuth(routes.ReadOneBlog)).Methods("GET")
	r.HandleFunc("/admin/blog/new", middleware.CheckAuth(routes.CreateBlog)).Methods("POST")
	r.HandleFunc("/admin/blog/edit/{id}", middleware.CheckAuth(routes.EditBlog)).Methods("PUT")
	r.HandleFunc("/admin/blog/delete/{id}", middleware.CheckAuth(routes.DeleteBlog)).Methods("DELETE")

	// Admin, Project Handlers
	r.HandleFunc("/admin/project/all", middleware.CheckAuth(routes.ReadProjects)).Methods("GET")
	r.HandleFunc("/admin/project/{id}", middleware.CheckAuth(routes.ReadOneProject)).Methods("GET")
	r.HandleFunc("/admin/project/new", middleware.CheckAuth(routes.CreateProject)).Methods("POST")
	r.HandleFunc("/admin/project/edit/{id}", middleware.CheckAuth(routes.EditProject)).Methods("PUT")
	r.HandleFunc("/admin/project/delete/{id}", middleware.CheckAuth(routes.DeleteProject)).Methods("DELETE")

	// Public, Document Handlers
	// This request get's redirected to the corresponding static file
	// If exists in static directory, then there, else source on web (Probably Gitlab/Github)
	r.HandleFunc("/public/blog/doc/{id}", routes.GetBlogDocument).Methods("GET")
	r.HandleFunc("/public/project/doc/{id}", routes.GetProjectDocument).Methods("GET")

	r.HandleFunc("/public/blog/all", routes.ReadPublicBlogs).Methods("GET")
	r.HandleFunc("/public/project/all", routes.ReadPublicProjects).Methods("GET")

	// For handling CORS requests
	ch := cors.New(cors.Options{
		AllowedOrigins:   config.Secrets.AllowedCorsURLs,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(r)

	// Server
	log.Fatal(http.ListenAndServe(":8080", ch))
}

// Dummy Index Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"hello world\", \"from\": \"Ritwik :)\"}"))
}
