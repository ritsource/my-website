package main

import (
	"fmt"
	"net/http"

	"github.com/ritwik310/my-website/server/handlers"
	"github.com/ritwik310/my-website/server/middleware"
)

func main() {
	fmt.Println("Hello world!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the website!")
	})

	http.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", middleware.CheckAuth(handlers.CurrentUser))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
