package mongo

import (
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type adminModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `json:"email"`
	GoogleID string        `json:"googleid"`
}

func newAdminModel(a *Admin) *adminModel {
	return &adminModel{
		Email: a.Email,
		GoogleID: a.GoogleID,
	}
}

func (a *adminModel) toAdmin() *Admin {
	return &Admin{
		ID:    a.ID.Hex(),
		Email: a.Email,
		GoogleID: a.GoogleID,
	}
}
