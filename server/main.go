package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// "github.com/julienschmidt/httprouter"
// "go.mongodb.org/mongo-driver/mongo"

var (
	isDev             bool    // Is in development mode
	mySecrets         secrets.Secrets // mySecrets
	googleOauthConfig *oauth2.Config
)

func init() {
	getSecrets(isDev, &mySecrets)

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     mySecrets.GoogleClientID,
		ClientSecret: mySecrets.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
	fmt.Println("isDev", isDev)

	fmt.Printf("%+v\n", mySecrets)

	http.HandleFunc("/auth/google", handleGoogleLogin)
	http.HandleFunc("/auth/google/callback", handleGoogleCallback)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Utils
// func Message(status bool, message string) map[string]interface{} {
// 	return map[string]interface{}{"status": status, "message": message}
// }

// func Respond(w http.ResponseWriter, data map[string]interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }

// Response Haldelers
func handler(w http.ResponseWriter, r *http.Request) {
	// tokenHeader := r.Header.Get("Authorization")
	fmt.Fprintf(w, "Hi there, Home here!")
}

var oauthStateString = "pseudo-random"

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
