package db

import (
	"flag"
	"os"
)

// GetDBName .
func GetDBName() string {
	if flag.Lookup("test.v") == nil {
		return os.Getenv("DB_NAME")
	}
	return os.Getenv("TEST_DB_NAME")
}

// Credenticals .
var (
	MongoURI = os.Getenv("MONGO_URI")
	DBName   = GetDBName()
)
