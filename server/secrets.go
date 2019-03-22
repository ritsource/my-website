package main

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
}

// Get Secrets for App
func getSecrets(isDev bool, mySecrets *Secrets) {
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
