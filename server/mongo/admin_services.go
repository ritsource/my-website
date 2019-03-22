package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AdminService ...
type AdminService struct {
	collection *mgo.Collection
}

// NewAdminService ...
func NewAdminService(session *Session, dbName string, collectionName string) *AdminService {
	collection := session.GetCollection(dbName, collectionName)
	fmt.Println("collection ==", collection)
	return &AdminService{collection: collection}
}

// Create ...
func (as *AdminService) Create(a *Admin) error {
	admin := newAdminModel(a)
	return as.collection.Insert(&admin)
}

// GetByEmail ...
func (as *AdminService) GetByEmail(email string) (*Admin, error) {
	model := adminModel{}
	err := as.collection.Find(bson.M{"email": email}).One(&model)
	return model.toAdmin(), err
}
