package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/mongo"
)

var (
	isDev               bool           // Is in development mode
	// Secrets ...
	Secrets             config.Secrets // mySecrets
	googleOauthConfig   *oauth2.Config
	mongoURL            string
	dbName              string
	adminCollectionName string
)

func init() {
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

var oauthStateString = "pseudo-random" // Oauth State

// GetHandleCurrentUser ...
func GetHandleCurrentUser(ms *mongo.Session) func(http.ResponseWriter, *http.Request) {
	// Real Handeler as return
	return func (w http.ResponseWriter, r *http.Request) {
		// Retrieving cookies from requests
		var err error
		var eCookie *http.Cookie // "admin-email" Cookie
		var hCookie *http.Cookie // "admin-id" Cookie

		eCookie, err = r.Cookie("admin-email")
		hCookie, err = r.Cookie("admin-id")

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
		admin, err = as.GetByEmail(eCookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Error: could't find admin"))
		}
		
		// Compare Hashed ID
		err = bcrypt.CompareHashAndPassword([]byte(hCookie.Value), []byte(admin.ID))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error: cookie didn't match"))
			fmt.Println(err)
		}
		
		// Marshaling admin
		var bData []byte // []byte - Admin Data
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
	admin, err := CreateOrGetAdmin(content)
	if err != nil {
		fmt.Printf("Error: query couldn't be done %v", err)
		return
	}

	// Generating hashed Cookie
	hCookie, err := GenSessionHash(SessionHashData{Email: admin.Email, ID: admin.ID})
	// Email Cookie - saves admin Email in the client cookie
	eCookie := http.Cookie{
		Name: "admin-email",
		Value: admin.Email,
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Path: "/",
		Domain: Secrets.DomainName,
	}
	
	// Setting the Cookie
	http.SetCookie(w, &hCookie) // Sets Hashed ID, in Cookie
	http.SetCookie(w, &eCookie) // Sets Email Cookie
	
	// Sending sesponse
	http.Redirect(w, r, "/auth/current_user", http.StatusTemporaryRedirect)

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
