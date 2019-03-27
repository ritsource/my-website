package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	gContext "github.com/gorilla/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/db"
)

var (
	isDev               bool // Is in development mode
	googleOauthConfig   *oauth2.Config
	mongoURL            string
	dbName              string
	adminCollectionName string
)

func init() {
	mongoURL = config.Secrets.MongoURI
	dbName = config.Secrets.DatabaseName

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     config.Secrets.GoogleClientID,
		ClientSecret: config.Secrets.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// GoogleLoginHandeler - ...
func GoogleLoginHandeler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GoogleCallbackHandler - ...
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// Get user info
	var content []byte
	err = nil
	content, err = GetUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: couldn't marshal data"))
		fmt.Println(err.Error())
		return
	}

	// Get admin if existes or Create new
	err = nil
	admin, err := CreateOrGetAdmin(content, db.Client)
	if err != nil {
		// w.Write([]byte("Error: query couldn't be done"))
		fmt.Println("Error: query couldn't be done")
	}

	// fmt.Println("admin.Email", admin.Email )
	// fmt.Println("admin.ID", admin.ID)

	// Generating hashed Cookie
	var hCookie http.Cookie
	err = nil
	hCookie, err = GenSessionHash(admin.GoogleID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: couldn't hash id"))
	}

	// Email Cookie - saves admin Email in the client cookie
	eCookie := http.Cookie{
		Name:    "admin-email",
		Value:   admin.Email,
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Path:    "/",
		Domain:  config.Secrets.DomainName,
	}

	// Setting the Cookie
	http.SetCookie(w, &hCookie) // Sets Hashed ID, in Cookie
	http.SetCookie(w, &eCookie) // Sets Email Cookie

	// Sending sesponse
	isDev = os.Getenv("isDev") == "true"
	if isDev {
		http.Redirect(w, r, config.Secrets.ConsoleCLientURL, http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "/auth/current_user", http.StatusTemporaryRedirect)
	}

	fmt.Println("Admin login successful!")
}

// CurrentUserHandler - checks currently logged in user
func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	admin := gContext.Get(r, "admin")

	// Marshaling admin
	bData, err := json.Marshal(admin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Error: couldn't marshal data"))
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bData)
}
