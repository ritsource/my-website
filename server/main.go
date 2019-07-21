package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ritwik310/my-website/server/handlers"
	mid "github.com/ritwik310/my-website/server/middleware"
	"github.com/ritwik310/my-website/server/renderers"
	"github.com/sirupsen/logrus"
)

func main() {
	// to clear the cache directory once at an interval
	if os.Getenv("DEV_MODE") != "true" {
		go func() {
			for {
				time.Sleep(2 * 24 * 60 * 60 * time.Second) // wait for two days after clearing
				err := os.RemoveAll("./cache/documents")   // haha
				if err != nil {
					logrus.Errorf("%v\n", err)
				}
				logrus.Infof("cleared cached files: %v", time.Now())
			}
		}()
	}

	// let's say
	http.HandleFunc("/", renderers.IndexHandler)
	http.HandleFunc("/blogs", renderers.BlogsHandler)
	http.HandleFunc("/blog/", renderers.BlogHandler)
	http.HandleFunc("/thread/", renderers.ThreadHandler)
	http.HandleFunc("/projects", renderers.ProjectsHandler)
	http.HandleFunc("/project/", renderers.ProjectHandler)
	http.HandleFunc("/preview", renderers.PreviewHandler)

	http.HandleFunc("/api/auth/google", handlers.GoogleLogin)
	http.HandleFunc("/api/auth/google/callback", handlers.GoogleCallback)
	http.HandleFunc("/api/auth/current_user", mid.CheckAuth(handlers.CurrentUser))

	http.HandleFunc("/api/private/blog", mid.CheckAuth(handlers.ReadBlog))
	http.HandleFunc("/api/private/blogs", mid.CheckAuth(handlers.ReadBlogs))
	http.HandleFunc("/api/private/blog/new", mid.CheckAuth(handlers.CreateBlog))
	http.HandleFunc("/api/private/blog/edit", mid.CheckAuth(handlers.EditBlog))
	http.HandleFunc("/api/private/blog/delete", mid.CheckAuth(handlers.DeleteBlog))
	http.HandleFunc("/api/private/blog/delete/permanent", mid.CheckAuth(handlers.DeleteBlogPrem))

	http.HandleFunc("/api/private/project", mid.CheckAuth(handlers.ReadProject))
	http.HandleFunc("/api/private/projects", mid.CheckAuth(handlers.ReadProjects))
	http.HandleFunc("/api/private/project/new", mid.CheckAuth(handlers.CreateProject))
	http.HandleFunc("/api/private/project/edit", mid.CheckAuth(handlers.EditProject))
	http.HandleFunc("/api/private/project/delete", mid.CheckAuth(handlers.DeleteProject))
	http.HandleFunc("/api/private/project/delete/permanent", mid.CheckAuth(handlers.DeleteProjectPrem))

	http.HandleFunc("/api/private/clear_cache", mid.CheckAuth(handlers.ClearCacheHandler))

	rfs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", rfs))

	http.ListenAndServe(":8080", nil)
}
