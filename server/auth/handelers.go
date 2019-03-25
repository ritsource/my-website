package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/mongo"
)

var (
	isDev bool // Is in development mode
	// Secrets ...
	Secrets             config.Secrets // mySecrets
	googleOauthConfig   *oauth2.Config
	mongoURL            string
	dbName              string
	adminCollectionName string
)

func init() {
	// Checking if in Development mode or not
	isDev = os.Getenv("isDev") == "true"
	fmt.Println("isDev:", isDev)

	// Getting env configs
	config.GetSecrets(isDev, &Secrets)

	mongoURL = Secrets.MongoURI
	dbName = Secrets.DatabaseName
	adminCollectionName = "admin"

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     Secrets.GoogleClientID,
		ClientSecret: Secrets.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

// GetCurrentUserHandeler ...
func GetCurrentUserHandeler(ms *mongo.Session) func(http.ResponseWriter, *http.Request) {
	// Real Handeler as return
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieving cookies from requests
		var err error
		var eCookie *http.Cookie // "admin-email" Cookie
		var hCookie *http.Cookie // "admin-id" Cookie

		eCookie, err = r.Cookie("admin-email")
		hCookie, err = r.Cookie("admin-id")

		err = nil
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error: admin unauthorized"))
			return
		}

		fmt.Println("eCookie.Value:", eCookie.Value)
		fmt.Println("hCookie.Value:", hCookie.Value)

		// as - Admin Service
		as := mongo.NewAdminService(ms.Copy(), dbName, adminCollectionName)

		// Query User from Database (by Email)
		var admin *mongo.Admin
		err = nil
		admin, err = as.GetByEmail(eCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Error: could't find admin"))
			return
		}

		// Compare Hashed ID
		err = nil
		err = bcrypt.CompareHashAndPassword([]byte(hCookie.Value), []byte(admin.ID))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error: cookie didn't match " + err.Error()))
			fmt.Println(err)
			return
		}

		// Marshaling admin
		var bData []byte // []byte - Admin Data
		err = nil
		bData, err = json.Marshal(admin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: couldn't marshal data"))
			fmt.Println(err)
		}

		w.Write(bData)

		// // http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		// w.Write([]byte("Hello World"))
	}
}

// HandleGoogleLogin ...
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GetGoogleCallbackHandeler ...
func GetGoogleCallbackHandeler(ms *mongo.Session) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		admin, err := CreateOrGetAdmin(content, ms)
		if err != nil {
			// w.Write([]byte("Error: query couldn't be done"))
			fmt.Println("Error: query couldn't be done")
		}

		// fmt.Println("admin.Email", admin.Email )
		// fmt.Println("admin.ID", admin.ID)

		// Generating hashed Cookie
		var hCookie http.Cookie
		err = nil
		hCookie, err = GenSessionHash(admin.ID)
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
			Domain:  Secrets.DomainName,
		}

		// Setting the Cookie
		http.SetCookie(w, &hCookie) // Sets Hashed ID, in Cookie
		http.SetCookie(w, &eCookie) // Sets Email Cookie

		// Sending sesponse
		if isDev {
			http.Redirect(w, r, Secrets.ConsoleCLientURL, http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}

		fmt.Println("Admin login successful..")
	}
}
