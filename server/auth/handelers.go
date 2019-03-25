package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ritwik310/my-website/server/config"
	"github.com/ritwik310/my-website/server/models"

	
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

// GetHandeler - ...
func GetHandeler(path string, client *mongo.Client) func(http.ResponseWriter, *http.Request) {
	var handeler func(http.ResponseWriter, *http.Request)

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
func googleLoginHandeler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// Returns "/auth/google/callback" handeler
func getGoogleCallback(client *mongo.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		fmt.Printf("r.FormValue %v , %v \n", r.FormValue("state"), r.FormValue("code"))

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
		isDev = os.Getenv("isDev") == "truef"
		if isDev {
			http.Redirect(w, r, config.Secrets.ConsoleCLientURL, http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "/auth/current_user", http.StatusTemporaryRedirect)
		}

		fmt.Println("Admin login successful..")
	}
}

// Returns "/auth/current_user" handeler
func getCurrentUser(client *mongo.Client) func(http.ResponseWriter, *http.Request) {
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

		// MongoDB collection
		collection := client.Database("dev_db").Collection("admins")

		// Query User from Database (by Email)
		var admin models.Admin
		err = nil
		err = collection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: eCookie.Value}}).Decode(&admin)		
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Error: could't find admin"))
			return
		}

		fmt.Printf("::::::::::::::::::::::: %+v", admin)

		// Compare Hashed ID
		err = nil
		err = bcrypt.CompareHashAndPassword([]byte(hCookie.Value), []byte(admin.GoogleID))
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
	}
}
