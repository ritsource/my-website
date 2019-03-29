package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// "github.com/ritwik310/my-website/api/auth"
	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/middleware"
	"github.com/ritwik310/my-website/api/routes"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "./static", "usage")
	flag.Parse()

	r := mux.NewRouter()

	// Static file server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/", indexHandler).Methods("GET")

	r.HandleFunc("/auth/google", routes.GoogleLoginHandeler).Methods("GET")
	r.HandleFunc("/auth/google/callback", routes.GoogleCallbackHandler).Methods("GET")
	r.HandleFunc("/auth/current_user", middleware.CheckAuth(routes.CurrentUserHandler)).Methods("GET")

	r.HandleFunc("/admin/blog/all", middleware.CheckAuth(routes.ReadBlogs)).Methods("GET")
	r.HandleFunc("/admin/blog/{id}", middleware.CheckAuth(routes.ReadOneBlog)).Methods("GET")
	r.HandleFunc("/admin/blog/new", middleware.CheckAuth(routes.CreateBlog)).Methods("POST")
	r.HandleFunc("/admin/blog/edit/{id}", middleware.CheckAuth(routes.EditBlog)).Methods("PUT")
	r.HandleFunc("/admin/blog/delete/{id}", middleware.CheckAuth(routes.DeleteBlog)).Methods("DELETE")

	r.HandleFunc("/admin/project/all", middleware.CheckAuth(routes.ReadProjects)).Methods("GET")
	r.HandleFunc("/admin/project/{id}", middleware.CheckAuth(routes.ReadOneProject)).Methods("GET")
	r.HandleFunc("/admin/project/new", middleware.CheckAuth(routes.CreateProject)).Methods("POST")
	r.HandleFunc("/admin/project/edit/{id}", middleware.CheckAuth(routes.EditProject)).Methods("PUT")
	r.HandleFunc("/admin/project/delete/{id}", middleware.CheckAuth(routes.DeleteProject)).Methods("DELETE")


	// r.HandleFunc("/public/blog/doc/{id}", middleware.CheckAuth(routes.GetBlogDocument)).	Methods("GET")
	r.HandleFunc("/public/project/doc/{id}", middleware.CheckAuth(routes.GetProjectDocument)).Methods("GET")

	// File Routes are Disables, cause there's no need of that, atleast not nowwww...
	// Down here!	
	// r.HandleFunc("/admin/file/all", routes.ReadFiles).Methods("GET")
	// r.HandleFunc("/admin/file/{id}", routes.ReadOneFile).Methods("GET")
	// r.HandleFunc("/admin/file/new", routes.CreateFile).Methods("POST")
	// r.HandleFunc("/admin/file/edit/{id}", routes.EditFile).Methods("PUT")
	// r.HandleFunc("/admin/file/delete/{id}", routes.DeleteFile).Methods("DELETE")

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
