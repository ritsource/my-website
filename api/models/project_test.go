package models_test

import (
	"testing"

	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	models.PCol.RemoveAll(nil)
}

func TestProjectCRUD(t *testing.T) {
	pr := models.Project{Title: "Test Project", Description: "Test Project"}

	createProject(t, &pr)
	readProject(t, &pr)
	if pr.IsDeleted == true {
		t.Error(".IsDeleted == true")
	}

	var prs models.Projects
	readProjects(t, &prs)
	if len(prs) < 1 {
		t.Error("len(prs) < 1")
	}

	// bl2 := models.Blog(bl)
	updateProject(t, &pr, bson.M{}, map[string]interface{}{"is_public": true})
	if pr.IsPublic == false {
		t.Error(".IsPublic == false")
	}

	deleteProject(t, &pr)
	if pr.IsDeleted == false {
		t.Error(".IsDeleted == false")
	}

	deleteProjectPermanent(t, &pr)
	readProjects(t, &prs)
	if len(prs) != 0 {
		t.Error("len(prs) > 0")
	}

}

func createProject(t *testing.T, pr *models.Project) {
	err := pr.Create()
	if err != nil {
		t.Error(err)
	}
}

func readProject(t *testing.T, pr *models.Project) {
	err := pr.Read(bson.M{}, bson.M{})
	if err != nil {
		t.Error(err)
	}
}

func readProjects(t *testing.T, prs *models.Projects) {
	err := prs.Read(bson.M{}, bson.M{})
	if err != nil {
		t.Error(err)
	}
}

func updateProject(t *testing.T, pr *models.Project, s, u bson.M) {
	err := pr.Update(s, u)
	if err != nil {
		t.Error(err)
	}
}

func deleteProject(t *testing.T, pr *models.Project) {
	err := pr.Delete(bson.ObjectId(""))
	if err != nil {
		t.Error(err)
	}
}

func deleteProjectPermanent(t *testing.T, pr *models.Project) {
	err := pr.DeletePermanent(bson.ObjectId(""))
	if err != nil {
		t.Error(err)
	}
}
