package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadResume .
func DownloadResume(w http.ResponseWriter, r *http.Request) {
	src := os.Getenv("RESUME_SOURCE")

	resp, err := http.Get(src)
	if err != nil {
		writeErr(w, 500, fmt.Errorf("internal server error"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Redirect(w, r, "/sorry", http.StatusSeeOther)
	}

	if r.URL.Query().Get("attachment") != "" {
		w.Header().Set("Content-Disposition", "attachment; filename=Resume.pdf")
	}

	w.Header().Set("Content-type", "application/pdf")

	io.Copy(w, resp.Body)
}
