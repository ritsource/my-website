package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

// ClearCacheAllHandler clears all the cached files
func ClearCacheAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	err := os.RemoveAll(path.Join(".", "cache", "documents")) // haha
	if err != nil {
		writeErr(w, 500, err)
	}
}

// ClearCacheSingleHandler delete a single document cache
func ClearCacheSingleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeErr(w, 404, fmt.Errorf("%v request to %v not found", r.Method, r.URL.Path))
		return
	}

	// retrieving id from query string
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("no object-id provided"))
		return
	}

	// deleting the id dir
	err := os.RemoveAll(path.Join(".", "cache", "documents", id))
	if err != nil {
		writeErr(w, 500, err)
	}
}
