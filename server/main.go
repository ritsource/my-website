package main

import (
	"fmt"
	"net/http"

	"github.com/ritwik310/my-website/server/handlers"
	mid "github.com/ritwik310/my-website/server/middleware"
	"github.com/ritwik310/my-website/server/renderers"
)

func main() {
	fmt.Println("Hello world!")

	http.HandleFunc("/", renderers.IndexHandler)
	http.HandleFunc("/blogs", renderers.BlogsHandler)
	http.HandleFunc("/blog/", renderers.BlogHandler)
	http.HandleFunc("/thread/", renderers.ThreadHandler)
	http.HandleFunc("/projects", renderers.ProjectsHandler)
	http.HandleFunc("/project/", renderers.ProjectHandler)

	http.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", mid.CheckAuth(handlers.CurrentUser))

	sfs := http.FileServer(http.Dir("raw/"))
	http.Handle("/raw/", http.StripPrefix("/raw/", sfs))

	rfs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", rfs))

	http.ListenAndServe(":8080", nil)
}
