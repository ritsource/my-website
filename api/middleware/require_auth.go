package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/ritwik310/my-website/api/config"
)

// Session - ...
var Session = sessions.NewCookieStore([]byte(config.Secrets.SessionKey))

// Writes Admin Un-Authenticated on Response
func writeUnauth(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("{\"message\": \"admin unauthorized\"}"))
}

// CheckAuth - Middleware that checks authentication
func CheckAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// checking Session
		session, _ := Session.Get(r, "session")
		aEmail, ok := session.Values["admin_id"]

		if !ok {
			fmt.Println("Error: authentication error")
			writeUnauth(w)
			return
		}

		if r.URL.Path == "/auth/current_user" {
			context.Set(r, "aEmail", aEmail)
		}

		handler.ServeHTTP(w, r)
	}
}
