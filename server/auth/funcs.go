package auth

import (
	"fmt"
	"errors"
	"encoding/json"
	"time"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	
	"github.com/ritwik310/my-website/server/mongo"
)

// Type for Data return from Google Oauth flow
type googleUserData struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// CreateOrGetAdmin ...
// Queries admin if it exists else Creates a new Admin in the Database 
func CreateOrGetAdmin(content []byte) (mongo.Admin, error) {
	session, connErr := mongo.NewSession(mongoURL)
	if connErr != nil {
		fmt.Printf("Error: unable to connect to mongo: %s\n", connErr)
	}

	// Close Connection
	defer session.Close()

	// Unmarshal Data returned by Google (content []byte)
	data, unmarErr := unmarshalAdmin(content)
	if unmarErr != nil {
		fmt.Printf("Error: unable to unmarshal %s\n", unmarErr)
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
		var nilAdmin mongo.Admin
		fmt.Printf("Email %s not Authorized..\n", data.Email)
		return nilAdmin, errors.New("Unauthorized")
	}

	// Admin Service
	adminService := mongo.NewAdminService(session.Copy(), dbName, adminCollectionName)

	// Check if user Exists in the Database
	admin, queryErr := adminService.Get(data.Email, data.ID)
	// Handle Error
	if queryErr != nil {
		fmt.Printf("Admin not found on Database %s\n", queryErr)
	} else {
		return *admin, nil
	}

	// Creating a new Admin if none Exist
	newAdmin := mongo.Admin{
		Email:    data.Email,
		GoogleID: data.ID,
	}

	insertErr := adminService.Create(&newAdmin)
	if insertErr != nil {
		fmt.Printf("Error: unable to insert new admin %s\n", insertErr)
		return newAdmin, insertErr
	}

	return newAdmin, nil
}

// Unmarshal Byte Slice to Struct
func unmarshalAdmin(content []byte) (googleUserData, error) {
	var admin googleUserData
	err := json.Unmarshal([]byte(content), &admin)
	return admin, err
}

// SessionHashData ...
type SessionHashData struct {
	Email string
	ID string
}

// GenSessionHash ...
func GenSessionHash(d SessionHashData) (http.Cookie, error) {
	var cookie http.Cookie // Cookie Struct
	
	// Generating Hash from byteData
	hashedData, hErr := bcrypt.GenerateFromPassword([]byte(d.ID), 14)
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
