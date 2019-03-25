package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Secrets ...
type Config struct {
	GoogleClientID     string `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `json:"GOOGLE_CLIENT_SECRET"`
	SessionKey         string `json:"SESSION_KEY"`
	MongoURI           string `json:"MONGO_URI"`
	DatabaseName       string `json:"DATABASE_NAME"`
	AdminEmails        []string `json:"ADMIN_EMAILS"`
	DomainName 				 string `json:"DOMAIN_NAME"`
	ConsoleCLientURL 	 string `json:"CONSOLE_CLIENT_URL"`
}

var isDev bool
// Secrets - ...
var Secrets Config

func init() {
	// Checking if in Development mode or not
	isDev = os.Getenv("isDev") == "true"
	fmt.Println("Development Mode:", isDev)

	// Getting env configs
	GetSecrets(isDev, &Secrets)
}

// GetSecrets - gets the secrets from Config.Dev file
func GetSecrets(isDev bool, s *Config) error {	
	var err error

	// JSON file location
	var filename string
	if isDev {
		filename = "config/config.development.json"
	} else {
		filename = "config/config.production.json"
	}

	// JSON file
	var jsonFile *os.File
	err = nil

	jsonFile, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	defer jsonFile.Close()

	// Reading JSON file
	var byteValue []byte
	err = nil

	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Saving data in struct
	err = nil
	err = json.Unmarshal([]byte(byteValue), &s)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
