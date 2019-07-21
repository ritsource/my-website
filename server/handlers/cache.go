package handlers

import (
	"net/http"
	"os"
)

// ClearCacheHandler clears all the cached files
func ClearCacheHandler(w http.ResponseWriter, r *http.Request) {
	err := os.RemoveAll("./cache/documents") // haha
	if err != nil {
		writeErr(w, 500, err)
	}
}
