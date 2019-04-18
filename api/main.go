package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/db"
	"github.com/ritwik310/my-website/api/middleware"
	"github.com/ritwik310/my-website/api/routes"
)

func main() {
	// Router (Gorilla Mux)
	r := mux.NewRouter()

	// Takes care of Static file serving
	var dir string
	flag.StringVar(&dir, "dir", "./cache", "usage")
	flag.Parse()

	// Static file server
	r.PathPrefix("/api/static/").Handler(http.StripPrefix("/api/static/", http.FileServer(http.Dir(dir))))

	// Dummy handler
	r.HandleFunc("/api", indexHandler).Methods("GET")

	// Authentication (Admin) handlers
	r.HandleFunc("/api/auth/google", routes.GoogleLoginHandeler).Methods("GET")
	r.HandleFunc("/api/auth/google/callback", routes.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/api/auth/current_user", middleware.CheckAuth(routes.CurrentUserHandler)).Methods("GET")

	// Admin, Blog Handlers
	r.HandleFunc("/api/private/blog/all", middleware.CheckAuth(routes.ReadBlogs)).Methods("GET")
	r.HandleFunc("/api/private/blog/{id}", middleware.CheckAuth(routes.ReadBlog)).Methods("GET")
	r.HandleFunc("/api/private/blog/new", middleware.CheckAuth(routes.CreateBlog)).Methods("POST")
	r.HandleFunc("/api/private/blog/edit/{id}", middleware.CheckAuth(routes.EditBlog)).Methods("PUT")
	r.HandleFunc("/api/private/blog/delete/{id}", middleware.CheckAuth(routes.DeleteBlog)).Methods("DELETE")
	r.HandleFunc("/api/private/project/delete/permanent/{id}", middleware.CheckAuth(routes.DeleteProjectF)).Methods("DELETE")

	// Admin, Project Handlers
	r.HandleFunc("/api/private/project/all", middleware.CheckAuth(routes.ReadProjects)).Methods("GET")
	r.HandleFunc("/api/private/project/{id}", middleware.CheckAuth(routes.ReadProject)).Methods("GET")
	r.HandleFunc("/api/private/project/new", middleware.CheckAuth(routes.CreateProject)).Methods("POST")
	r.HandleFunc("/api/private/project/edit/{id}", middleware.CheckAuth(routes.EditProject)).Methods("PUT")
	r.HandleFunc("/api/private/project/delete/{id}", middleware.CheckAuth(routes.DeleteProject)).Methods("DELETE")
	r.HandleFunc("/api/private/project/delete/permanent/{id}", middleware.CheckAuth(routes.DeleteBlogF)).Methods("DELETE")

	// Public, Document Handlers
	// This request get's redirected to the corresponding static file
	// If exists in static directory, then there, else source on web (Probably Gitlab/Github)
	r.HandleFunc("/api/public/blog/doc/{id}", routes.PubGetBlogDoc).Methods("GET")
	r.HandleFunc("/api/public/project/doc/{id}", routes.PubGetProjectDoc).Methods("GET")

	// Public API Routes
	r.HandleFunc("/api/public/blog/all", routes.PubReadBlogs).Methods("GET")
	r.HandleFunc("/api/public/blog/{id}", routes.PubReadBlog).Methods("GET")
	r.HandleFunc("/api/public/project/all", routes.PubReadProjects).Methods("GET")
	r.HandleFunc("/api/public/project/{id}", routes.PubReadProject).Methods("GET")

	// For handling CORS
	ch := cors.New(cors.Options{
		AllowedOrigins:   []string{config.Secrets.AppRendererURL, config.Secrets.ConsoleCLientURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(r)

	// Listening in PORT 8080
	log.Fatal(http.ListenAndServe(":8080", ch))

	// Closes Database, on main.go File Exit
	defer db.Close()
}

// Dummy Index Handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Hello World, this is API\", \"from\": \"Ritwik :)\"}"))
}
