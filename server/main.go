package main

import (
	// "fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"github.com/ritwik310/my-website/server/auth"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/handlers"
	"github.com/ritwik310/my-website/server/middleware"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	router.GET("/auth/current_user", auth.PickHandeler("/auth/current_user", db.Client))
	router.GET("/auth/google", auth.PickHandeler("/auth/google", db.Client))
	router.GET("/auth/google/callback", auth.PickHandeler("/auth/google/callback", db.Client))

	blog := handlers.Blog{
		Client: db.Client,
		Db:     "dev_db",
		Col:    "blogs",
	}

	// router.READ("/admin/blog/:id", )
	router.GET("/admin/blog/all", middleware.AuthRequired(blog.ReadAll))
	router.POST("/admin/add_blog", blog.CreateOne)
	// router.PUT("/admin/edit_blog/:id", )
	// router.DELETE("/admin/delete_blog/:id", )

	// log.Fatal(http.ListenAndServe(":8080", nil))
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	}).Handler(router)

	http.ListenAndServe(":8080", handler)
}
