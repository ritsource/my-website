package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	GoogleClientID     string   `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string   `json:"GOOGLE_CLIENT_SECRET"`
	SessionKey         string   `json:"SESSION_KEY"`
	MongoURI           string   `json:"MONGO_URI"`
	BackupMongoURI     string   `json:"BACKUP_MONGO_URI"`
	DBName             string   `json:"DATABASE_NAME"`
	AdminEmails        []string `json:"ADMIN_EMAILS"`
	ConsoleCLientURL   string   `json:"CONSOLE_CLIENT_URL"`
	AllowedCorsURLs    []string `json:"ALLOWED_CORS_URLS"`
}

var isDev bool

// Secrets - ...
var Secrets Config

func init() {
	// Checking if in Development mode or not
	isDev = os.Getenv("isDev") == "true"
	fmt.Println("Development Mode =", isDev)

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

// func getSecrets(s *Config) {
// 	s.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
// 	s.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
// 	s.SessionKey = os.Getenv("SESSION_KEY")
// 	s.MongoURI = os.Getenv("MONGO_URI")
// 	s.BackupMongoURI = os.Getenv("BACKUP_MONGO_URI")
// 	s.DBName = os.Getenv("DATABASE_NAME")
// 	s.AdminEmails = os.Getenv("ADMIN_EMAILS")
// 	s.ConsoleCLientURL = os.Getenv("CONSOLE_CLIENT_URL")
// 	s.AllowedCorsURLs = os.Getenv("isDev")
// }
