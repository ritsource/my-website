package models_test

import (
	"testing"

	"github.com/ritwik310/my-website/api/models"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	models.BCol.RemoveAll(nil)
}

func TestBlogCRUD(t *testing.T) {
	bl := models.Blog{Title: "Test Blog", Description: "Test Blog"}

	createBlog(t, &bl)
	readBlog(t, &bl)
	if bl.IsDeleted == true {
		t.Error(".IsDeleted == true")
	}

	var bls models.Blogs
	readBlogs(t, &bls)
	if len(bls) < 1 {
		t.Error("len(bls) < 1")
	}

	// bl2 := models.Blog(bl)
	updateBlog(t, &bl, bson.M{}, map[string]interface{}{"is_public": true})
	if bl.IsPublic == false {
		t.Error(".IsPublic == false")
	}

	deleteBlog(t, &bl)
	if bl.IsDeleted == false {
		t.Error(".IsDeleted == false")
	}

	deleteBlogPermanent(t, &bl)
	readBlogs(t, &bls)
	if len(bls) != 0 {
		t.Error("len(bls) > 0")
	}

}

func createBlog(t *testing.T, bl *models.Blog) {
	err := bl.Create()
	if err != nil {
		t.Error(err)
	}
}

func readBlog(t *testing.T, bl *models.Blog) {
	err := bl.Read(bson.M{}, bson.M{})
	if err != nil {
		t.Error(err)
	}
}

func readBlogs(t *testing.T, bls *models.Blogs) {
	err := bls.Read(bson.M{}, bson.M{})
	if err != nil {
		t.Error(err)
	}
}

func updateBlog(t *testing.T, bl *models.Blog, s, u bson.M) {
	err := bl.Update(s, u)
	if err != nil {
		t.Error(err)
	}
}

func deleteBlog(t *testing.T, bl *models.Blog) {
	err := bl.Delete(bson.ObjectId(""))
	if err != nil {
		t.Error(err)
	}
}

func deleteBlogPermanent(t *testing.T, bl *models.Blog) {
	err := bl.DeletePermanent(bson.ObjectId(""))
	if err != nil {
		t.Error(err)
	}
}
