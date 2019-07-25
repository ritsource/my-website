package main

import (
	"net/http"
	"os"

	"github.com/ritwik310/my-website/server/handlers"
	mid "github.com/ritwik310/my-website/server/middleware"
	"github.com/ritwik310/my-website/server/renderers"
	"github.com/rs/cors"
)

func main() {
	// to clear the cache directory once at an interval
	// if os.Getenv("DEV_MODE") != "true" {
	// 	go func() {
	// 		for {
	// 			time.Sleep(10 * 24 * 60 * 60 * time.Second) // wait for two days after clearing
	// 			err := os.RemoveAll("./cache/documents")    // haha
	// 			if err != nil {
	// 				logrus.Errorf("%v\n", err)
	// 			}
	// 			logrus.Infof("cleared cached files: %v", time.Now())
	// 		}
	// 	}()
	// }

	mux := http.NewServeMux()

	mux.HandleFunc("/", renderers.IndexHandler)
	mux.HandleFunc("/blogs", renderers.BlogsHandler)
	mux.HandleFunc("/blog/", renderers.BlogHandler)
	mux.HandleFunc("/thread/", renderers.ThreadHandler)
	mux.HandleFunc("/projects", renderers.ProjectsHandler)
	mux.HandleFunc("/project/", renderers.ProjectHandler)
	mux.HandleFunc("/preview", renderers.PreviewHandler)

	mux.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	mux.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	mux.HandleFunc("/api/auth/current_user", mid.CheckAuth(handlers.CurrentUser))

	mux.HandleFunc("/api/private/blog", mid.CheckAuth(handlers.ReadBlog))
	mux.HandleFunc("/api/private/blogs", mid.CheckAuth(handlers.ReadBlogs))
	mux.HandleFunc("/api/private/blog/new", mid.CheckAuth(handlers.CreateBlog))
	mux.HandleFunc("/api/private/blog/edit", mid.CheckAuth(handlers.EditBlog))
	mux.HandleFunc("/api/private/blog/delete", mid.CheckAuth(handlers.DeleteBlog))
	mux.HandleFunc("/api/private/blog/delete/permanent", mid.CheckAuth(handlers.DeleteBlogPrem))

	mux.HandleFunc("/api/private/project", mid.CheckAuth(handlers.ReadProject))
	mux.HandleFunc("/api/private/projects", mid.CheckAuth(handlers.ReadProjects))
	mux.HandleFunc("/api/private/project/new", mid.CheckAuth(handlers.CreateProject))
	mux.HandleFunc("/api/private/project/edit", mid.CheckAuth(handlers.EditProject))
	mux.HandleFunc("/api/private/project/delete", mid.CheckAuth(handlers.DeleteProject))
	mux.HandleFunc("/api/private/project/delete/permanent", mid.CheckAuth(handlers.DeleteProjectPrem))

	// TODO: cache enable with GKE cluster
	// mux.HandleFunc("/api/private/clear_cache/all", mid.CheckAuth(handlers.ClearCacheAllHandler))
	// mux.HandleFunc("/api/private/clear_cache", mid.CheckAuth(handlers.ClearCacheSingleHandler))

	rfs := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", rfs))

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ADMIN_ORIGIN")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)

	http.ListenAndServe(":8080", handler)
}
