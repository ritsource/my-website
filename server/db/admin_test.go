package db_test

import (
	"testing"

	"github.com/ritcrap/my-website/server/db"
	"gopkg.in/mgo.v2/bson"
)

func TestAdminCRUD(t *testing.T) {
	email := "ritwiktesting@example.com"
	googleId := "123456789010"

	a := db.Admin{
		Email:    email,
		GoogleID: googleId,
	}

	err := a.Create()
	if err != nil {
		t.Error(err)
	}

	if a.Email != email || a.GoogleID != googleId {
		t.Error("a.Email != email || a.GoogleID != googleId")
	}

	var a2 db.Admin
	err = a2.Read(bson.M{"email": email, "google_id": googleId})
	if err != nil {
		t.Error(err)
	}

	if a.Email != a2.Email || a.GoogleID != a2.GoogleID {
		t.Error("a.Email != a2.Email || a.GoogleID != a2.GoogleID")
	}
}
