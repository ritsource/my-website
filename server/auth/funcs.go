package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
	"golang.org/x/oauth2"
)

// Type for Data return from Google Oauth flow
type googleUserData struct {
	Email    string `json:"email"`
	GoogleID string `json:"id"`
}

var oauthStateString = "pseudo-random" // Oauth State

// GetUserInfo - gets User (Admin) info from Google APIs
func GetUserInfo(state string, code string) ([]byte, error) {
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