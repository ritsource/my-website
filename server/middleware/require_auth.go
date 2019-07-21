package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

// CookieStore .
var CookieStore *sessions.CookieStore

func init() {
	CookieStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

// CheckAuth .
func CheckAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := CookieStore.Get(r, "session")
		if err != nil {
			fmt.Println("x1")
			w.WriteHeader(500)
			w.Write([]byte("{\"message\": \"" + err.Error() + "\"}"))
			return
		}

		email, ok := session.Values["admin_email"]

		if !ok {
			fmt.Println("x2")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("{\"message\": \"not authorized - please login as admin\"}"))
			return
		}

		if r.URL.Path == "/api/auth/current_user" {
			context.Set(r, "admin_e", email)
		}

		handler.ServeHTTP(w, r)
	}
}
