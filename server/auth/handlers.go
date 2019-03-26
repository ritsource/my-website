package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/julienschmidt/httprouter"

	"github.com/ritwik310/my-website/server/config"

	
)

var (
	isDev bool // Is in development mode
	googleOauthConfig   *oauth2.Config
	mongoURL            string
	dbName              string
	adminCollectionName string
)

func init() {
	mongoURL = config.Secrets.MongoURI
	dbName = config.Secrets.DatabaseName
	adminCollectionName = "admin"

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     config.Secrets.GoogleClientID,
		ClientSecret: config.Secrets.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// PickHandeler - ...
func PickHandeler(path string, client *mongo.Client) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	var handeler func(http.ResponseWriter, *http.Request, httprouter.Params)

	switch path {
	case "/auth/google":
		handeler = googleLoginHandeler
	case "/auth/google/callback":
		handeler = getGoogleCallback(client)
	case "/auth/current_user":
		handeler = getCurrentUser(client)		
	}

	return handeler
}

// HandleGoogleLogin ...
func googleLoginHandeler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// Returns "/auth/google/callback" handeler
func getGoogleCallback(client *mongo.Client) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		admin, err := CreateOrGetAdmin(content, client)
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
}

// Returns "/auth/current_user" handeler
func getCurrentUser(client *mongo.Client) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//  Checking out current user
		admin, err := CheckAuth(r, client)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error:" + err.Error()))
			fmt.Println(err)
			return
		}

		// Marshaling admin
		var bData []byte // []byte - Admin Data
		err = nil
		bData, err = json.Marshal(admin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Error: couldn't marshal data"))
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(bData)
	}
}
