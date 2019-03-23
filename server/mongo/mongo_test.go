package mongo_test

import (
	"log"
	"testing"

	"github.com/ritwik310/my-website/server/mongo"
)

const (
	mongoUrl            = "localhost:27017"
	dbName              = "test_db"
	adminCollectionName = "admin"
)

func Test_AdminService(t *testing.T) {
	t.Run("CreateAdmin", insertAdmin)
}

func insertAdmin(t *testing.T) {
	// Arrange
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer session.Close()
	adminService := mongo.NewAdminService(session.Copy(), dbName, adminCollectionName)

	testEmail := "integrationTest@example.com"
	admin := mongo.Admin{
		Email: testEmail,
	}

	// Act
	err = adminService.Create(&admin)

	// Assert
	if err != nil {
		t.Errorf("Unable to create admin: %s", err)
	}
	var results []mongo.Admin
	session.GetCollection(dbName, adminCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Email != admin.Email {
		t.Errorf("Incorrect email. Expected `%s`, Got: `%s`", testEmail, results[0].Email)
	}
}
