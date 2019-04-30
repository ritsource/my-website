package handlers

import (
	"net/http"
)

// NotFoundHandler ...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	RenderError(w, 404, "Page Not Found")
}
