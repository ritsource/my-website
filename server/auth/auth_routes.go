package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"

	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/models"
)

// Session - ...
var Session = sessions.NewCookieStore([]byte("config.Secrets.SessionKey"))

// Writes Admin Un-Authenticated on Response
func writeError(w http.ResponseWriter, status int, err error, msg string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\"}"))
	fmt.Println("Error:", err.Error())
}

// GoogleOauthConfig - ...
var GoogleOauthConfig *oauth2.Config

func init() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     config.Secrets.GoogleClientID,
		ClientSecret: config.Secrets.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// GoogleLoginHandeler - ...
func GoogleLoginHandeler(w http.ResponseWriter, r *http.Request) {
	url := GoogleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GoogleCallbackHandler - ...
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// Get user info in []byte
	var content []byte
	content, err = GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		writeError(w, 500, err, "Couldn't read google's data")
		return
	}

	// Unmarshal Data
	var data models.Admin
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		writeError(w, 500, err, "Couldn't read google's data")
		return
	}

	var admin models.Admin

	// Query Admin
	admin, err = admin.Read(bson.M{"email": data.Email, "google_id": data.GoogleID})
	if err != nil {
		// Inserting Document
		admin, err = data.Create()
		if err != nil {
			writeError(w, 422, err, "Failed to insert new document")
			return
		}
	}

	// Saving Session
	session, err := Session.Get(r, "session")
	session.Values["admin_id"] = admin.Email
	session.Save(r, w)

	// Sending sesponse
	isDev := os.Getenv("isDev") == "true"
	fmt.Println("isDev", isDev)
	if isDev {
		http.Redirect(w, r, config.Secrets.ConsoleCLientURL, http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "/auth/current_user", http.StatusTemporaryRedirect)
	}

	fmt.Println("Login successful!")
}

// CurrentUserHandler - checks currently logged in user
func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	// var aEmail string
	aEmail := context.Get(r, "aEmail")

	var admin models.Admin
	var err error

	// Query Admin
	admin, err = admin.Read(bson.M{"email": aEmail})
	if err != nil {
		writeError(w, 422, err, "Unable to fild admin in db")
		return
	}

	bData, err := json.Marshal(admin)
	if err != nil {
		writeError(w, 500, err, "Error: couldn't marshal data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}
