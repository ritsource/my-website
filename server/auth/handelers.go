package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/ritwik310/my-website/server/config"
)

var (
	isDev               bool           // Is in development mode
	secrets             config.Secrets // mySecrets
	googleOauthConfig   *oauth2.Config
	mongoURL            string
	dbName              string
	adminCollectionName string
)

func init() {
	config.GetSecrets(isDev, &secrets)

	mongoURL = secrets.MongoURI
	// dbName = secrets.DatabaseName
	dbName = "test_db"
	adminCollectionName = "admin"

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     secrets.GoogleClientID,
		ClientSecret: secrets.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

var oauthStateString = "pseudo-random" // Oauth State

// HandleCurrentUser ...
func HandleCurrentUser(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// HandleGoogleLogin ...
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// HandleGoogleCallback ...
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Get user info
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// fmt.Printf("%v", string(content))
	CreateOrGetUser(content)
	// Encrypt
	w.Write(content)

	// http.Redirect(w, r, "/", http.StatusMovedPermanently) // Redirecting to Home
}

// Gets User (Admin) info from Google APIs
func getUserInfo(state string, code string) ([]byte, error) {
	// Checking Oauth-State
	if state != oauthStateString {
		return nil, fmt.Errorf("Invalid oauth state")
	}

	// Getting Token
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
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
