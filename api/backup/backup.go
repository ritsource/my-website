package backup

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/ritwik310/my-website/api/models"

	"github.com/ritwik310/my-website/api/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Local Session
var lClient *mgo.Session

// Backup Session
var bClient *mgo.Session

func init() {
	var err error

	lClient, err = mgo.Dial(config.Secrets.MongoURI)
	bClient, err = mgo.Dial(config.Secrets.BackupMongoURI)

	if err != nil {
		fmt.Println("Error: couldn't connect to MongoDB")
		log.Fatal(err)
	}
}

// BackItup -
func BackItup() {
	for {
		// Run Only once every day
		time.Sleep(time.Hour * 24)

		ch := make(chan string)

		go cloneBlogs(ch)
		go cloneProjects(ch)
	}
}

func cloneBlogs(ch chan string) {
	var blogs models.Blogs

	blogs, err := blogs.Read(bson.M{})
	if err != nil {
		// return nil, err
		return
	}

	// Inserting to the backup Database
	c := make(chan string)

	fmt.Println("Inserting Blogs to the backup Database..")

	for _, b := range blogs {
		// Inserting values to the backup DB
		go insertDoc("blogs", b, c)
	}

	ch <- "Blogs Done!"
}

func cloneProjects(ch chan string) {
	var projects models.Projects

	// Getting Documents from Master
	projects, err := projects.Read(bson.M{})
	if err != nil {
		fmt.Println("Error: couldn't read project from main DB", err)
		return
	}

	// Inserting to the backup Database
	c := make(chan string)

	fmt.Println("Inserting Projects to the backup Database..")

	for _, p := range projects {
		// Inserting values to the backup DB
		go insertDoc("projects", p, c)
	}

	ch <- "Projects Done!"
}

func insertDoc(colName string, val interface{}, c chan string) {
	err := bClient.DB(config.Secrets.DBName).C(colName).Insert(val)
	c <- "Inserted!"

	if err != nil {
		dup, _ := regexp.Match(err.Error(), []byte("E11000"))
		if dup {
			fmt.Println("Already Exists!")
			return
		}

		fmt.Println("Error: couldn't insert to the backup database", err)
	}
}
