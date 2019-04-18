package config

import (
	"os"
)

// Config ...
type Config struct {
	GoogleClientID     string   `json:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string   `json:"GOOGLE_CLIENT_SECRET"`
	SessionKey         string   `json:"SESSION_KEY"`
	MongoURI           string   `json:"MONGO_URI"`
	DBName             string   `json:"DATABASE_NAME"`
	AdminEmails        []string `json:"ADMIN_EMAILS"`
	ConsoleCLientURL   string   `json:"CONSOLE_CLIENT_URL"`
	AppRendererURL     string   `json:"APP_RENDERER_URL"`
}

// Secrets - Struct (Config) that holds environment variables values
var Secrets Config
var isDev bool

func init() {
	isDev = os.Getenv("DEV_MODE") == "true" // Checking if in Development mode or not
	ReadSecrets(&Secrets)                   // Reading env configs
}

// ReadSecrets - Gets secrets from environment variables
func ReadSecrets(s *Config) {
	s.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	s.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	s.SessionKey = os.Getenv("SESSION_KEY")
	s.MongoURI = os.Getenv("MONGO_URI")
	s.DBName = os.Getenv("DATABASE_NAME")

	AdminEmailA := os.Getenv("ADMIN_EMAIL_A")
	AdminEmailB := os.Getenv("ADMIN_EMAIL_B")
	s.AdminEmails = append(s.AdminEmails, AdminEmailA, AdminEmailB)

	s.ConsoleCLientURL = os.Getenv("CONSOLE_CLIENT_URL")
	s.AppRendererURL = os.Getenv("APP_RENDERER_URL")
}
