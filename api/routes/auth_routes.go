package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/mgo.v2/bson"

	"github.com/ritwik310/my-website/api/config"
	"github.com/ritwik310/my-website/api/middleware"
	"github.com/ritwik310/my-website/api/models"
)

// oauthStateString
var oauthStateString = "pseudo-random"

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

// WriteError Admin Un-Authenticated on Response
func WriteError(w http.ResponseWriter, status int, err error, msg string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"" + msg + "\"}"))
	fmt.Println("Error:", err.Error())
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
	content, err = getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		WriteError(w, 500, err, "Couldn't read google's data")
		return
	}

	// Unmarshal Data
	var data models.Admin
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		WriteError(w, 500, err, "Couldn't read google's data")
		return
	}

	var admin models.Admin

	// Query Admin
	admin, err = admin.Read(bson.M{"email": data.Email, "google_id": data.GoogleID})
	if err != nil {
		// Inserting Document
		admin, err = data.Create()
		if err != nil {
			WriteError(w, 422, err, "Failed to insert new document")
			return
		}
	}

	// Saving Session
	session, err := middleware.Session.Get(r, "session")
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
		WriteError(w, 422, err, "Unable to fild admin in db")
		return
	}

	bData, err := json.Marshal(admin)
	if err != nil {
		WriteError(w, 500, err, "Error: couldn't marshal data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}

// Helper Functions ...
// getUserInfo - gets User (Admin) info from Google APIs
func getUserInfo(state string, code string) ([]byte, error) {
	// Checking Oauth-State
	if state != oauthStateString {
		return nil, fmt.Errorf("Invalid oauth state")
	}

	// Getting Token
	token, err := GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("Code exchange failed: %s", err.Error())
	}

	// Getting User info
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return content, nil
}
