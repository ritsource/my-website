package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Secrets ...
type Secrets struct {
	GoogleClientID     string `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `json:"GOOGLE_CLIENT_SECRET"`
	SessionKey         string `json:"SESSION_KEY"`
	MongoURI           string `json:"MONGO_URI"`
	DatabaseName       string `json:"DATABASE_NAME"`
}

// GetSecrets ...
func GetSecrets(isDev bool, mySecrets *Secrets) {
	isDev = os.Getenv("isDev") == "true"

	var filename string // JSON file location

	if isDev {
		filename = "config/config.development.json"
	} else {
		filename = "config/config.production.json"
	}

	jsonFile, readErr := os.Open(filename)

	if readErr != nil {
		fmt.Println("Error", readErr)
	}

	fmt.Println("Successfully Opened config json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	marshErr := json.Unmarshal([]byte(byteValue), &mySecrets)

	if marshErr != nil {
		fmt.Println("Error", marshErr)
	}
}
