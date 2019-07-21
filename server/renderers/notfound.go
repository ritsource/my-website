package renderers

import (
	"net/http"
)

// NotFoundHandler ...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderErr(w, 404, "Page Not Found")
}
