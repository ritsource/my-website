package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/crypto/bcrypt"

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
	dbName = secrets.DatabaseName
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
	
	// Get admin if existes or Create new
	admin, err := CreateOrGetUser(content)
	if err != nil {
		fmt.Printf("Error: query couldn't be done %v", err)
		return
	}

	// Turn Admin into []byte
	byteData, marshalErr := json.Marshal(admin)
	if marshalErr != nil {
		fmt.Printf("Error: couldn't marshal admin data %v", marshalErr)
		return
	}
	
	// Generating Hash from byteData
	hashedData, err := bcrypt.GenerateFromPassword(byteData, 14)

	// Setting Hashed Data in Cookie
	cookie := http.Cookie{
		Name: "session",
		Value: string(hashedData),
		Expires: time.Now().Add(30 * 24 * time.Hour),
	}
	
	// Setting the Cookie
	http.SetCookie(w, &cookie)
	
	// Sending sesponse
	w.Write(byteData)

	fmt.Println("Admin login successful..")
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
