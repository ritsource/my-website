package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"

	"github.com/ritwik310/my-website/server/mongo"
)

// Type for Data return from Google Oauth flow
type googleUserData struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

var oauthStateString = "pseudo-random" // Oauth State

// GetUserInfo - gets User (Admin) info from Google APIs
func GetUserInfo(state string, code string) ([]byte, error) {
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

// CreateOrGetAdmin - queries Admin if it exists else Creates a new Admin in the Database
func CreateOrGetAdmin(content []byte, ms *mongo.Session) (mongo.Admin, error) {
	var err error
	var admin *mongo.Admin

	// Unmarshal Data returned by Google (content []byte)
	var data googleUserData
	err = nil
	data, err = unmarshalAdmin(content)
	if err != nil {
		fmt.Printf("Error: unable to unmarshal %s\n", err)
		return *admin, err
	}

	// Checking if Email Allowed or Not
	userUnauth := true // User Unauthorized
	for _, email := range Secrets.AdminEmails {
		if email == data.Email {
			userUnauth = false
		}
	}
	// To handle Unauthorized Players
	if userUnauth {
		fmt.Printf("Email %s not Authorized..\n", data.Email)
		return *admin, errors.New("Unauthorized")
	}

	// as - Admin Service
	as := mongo.NewAdminService(ms.Copy(), dbName, adminCollectionName)

	// Check if user Exists in the Database
	err = nil
	admin, err = as.Get(data.Email, data.ID)
	if err != nil {
		fmt.Printf("Admin not found on Database %s\n", err)
	} else {
		return *admin, err
	}

	// Creating a new Admin if none Exist
	newAdmin := mongo.Admin{
		Email:    data.Email,
		GoogleID: data.ID,
	}

	// Creating Mongo Document
	err = nil
	err = as.Create(&newAdmin)
	if err != nil {
		fmt.Printf("Error: unable to insert new admin %s\n", err)
		return newAdmin, err
	}

	fmt.Printf("Here: newAdmin %+v\n", newAdmin)

	return newAdmin, nil
}

// Unmarshal Byte Slice to Struct
func unmarshalAdmin(content []byte) (googleUserData, error) {
	var admin googleUserData
	err := json.Unmarshal([]byte(content), &admin)
	return admin, err
}

// GenSessionHash ...
func GenSessionHash(id string) (http.Cookie, error) {
	var cookie http.Cookie // Cookie Struct
	fmt.Println("var cookie http.Cookie "+ id)

	// Generating Hash from byteData
	hashedData, hErr := bcrypt.GenerateFromPassword([]byte(id), 14)
	if hErr != nil {
		return cookie, hErr
	}

	cookie.Name = "admin-id"
	cookie.Value = string(hashedData)
	cookie.Expires = time.Now().Add(30 * 24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = Secrets.DomainName

	return cookie, nil
}
