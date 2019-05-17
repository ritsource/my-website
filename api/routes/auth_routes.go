package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		RedirectURL:  config.Secrets.GoogleAuthRedirectURL,
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

// GoogleLoginHandeler - Handles request for /auth/google
func GoogleLoginHandeler(w http.ResponseWriter, r *http.Request) {
	url := GoogleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GoogleCallbackHandler - Handles Google Callback, /auth/google/callback
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get user info in []byte
	var content []byte
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	HandleErr(w, 500, err)

	// Unmarshal Data
	var data models.Admin
	err = json.Unmarshal([]byte(content), &data)
	HandleErr(w, 500, err)

	var admin models.Admin

	// Query Admin
	err = admin.Read(bson.M{"email": data.Email, "google_id": data.GoogleID})
	if err != nil {
		// Inserting Document
		admin = data
		err = admin.Create()
		HandleErr(w, 500, err)
	}

	// Saving Session
	session, err := middleware.Session.Get(r, "session")
	session.Values["admin_id"] = admin.Email
	session.Save(r, w)

	// Sending sesponse
	// isDev := os.Getenv("DEV_MODE") == "true"
	// if isDev {
	http.Redirect(w, r, config.Secrets.ConsoleClientURL+"/admin", http.StatusTemporaryRedirect)
	// } else {
	// 	http.Redirect(w, r, "/api/auth/current_user", http.StatusTemporaryRedirect)
	// }

	fmt.Println("Login successful!")
}

// CurrentUserHandler - checks currently logged in user
func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	// var aEmail string
	aEmail := context.Get(r, "aEmail")
	var admin models.Admin

	// Query Admin
	err := admin.Read(bson.M{"email": aEmail})
	HandleErr(w, 422, err)

	bData, err := json.Marshal(admin)
	HandleErr(w, 500, err)

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
