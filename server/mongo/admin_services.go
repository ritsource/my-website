package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

// "github.com/ritwik310/my-website/server/auth"

// AdminService ...
type AdminService struct {
	collection *mgo.Collection
}

// NewAdminService ...
func NewAdminService(session *Session, dbName string, collectionName string) *AdminService {
	collection := session.GetCollection(dbName, collectionName)
	return &AdminService{collection: collection}
}

// Create ...
func (as *AdminService) Create(a *Admin) error {
	admin := newAdminModel(a)
	return as.collection.Insert(&admin)
}

// Get ...
func (as *AdminService) Get(Email string, ID string) (*Admin, error) {
	model := adminModel{}
	err := as.collection.Find(bson.M{"email": Email, "googleid": ID}).One(&model)
	return model.toAdmin(), err
}

// GetByEmail ...
func (as *AdminService) GetByEmail(Email string) (*Admin, error) {
	model := adminModel{}
	err := as.collection.Find(bson.M{"email": Email}).One(&model)
	fmt.Printf("%+v/n", model.toAdmin())
	return model.toAdmin(), err
}
