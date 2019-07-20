package main

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ritwik310/my-website/server/handlers"
	mid "github.com/ritwik310/my-website/server/middleware"
	"github.com/ritwik310/my-website/server/renderers"
)

func main() {
	// go ClearCache()

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

// ClearCache clears the cache directory once at an interval
func ClearCache() {
	for {
		err := os.RemoveAll("./cache/documents")
		if err != nil {
			logrus.Errorf("%v\n", err)
		}

		// haha
		logrus.Infof("cleared cached files: %v", time.Now())

		// wait for two days after clearing
		time.Sleep(2 * 24 * 60 * 60 * time.Second)
	}
}
