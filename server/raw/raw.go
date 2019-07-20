package raw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/ritwik310/my-website/server/db"
	"github.com/sirupsen/logrus"
)

// GetDocument .
func GetDocument(id string, src string, doctype int8, index int) ([]byte, error) {

	// constructing cahce filename and getting the actual source file url
	var fp string
	switch doctype {
	case db.DocTypeMD:
		fp = path.Join(".", "cache", id, strconv.Itoa(index)+".md")
	case db.DocTypeHTML:
		fp = path.Join(".", "cache", id, strconv.Itoa(index)+".html")
	default:
		return nil, fmt.Errorf("document type not defined")
	}

	// checking if requested file Exists or Not
	if _, err := os.Stat(fp); err == nil {
		return ioutil.ReadFile(fp)
	}

	// reading data from remote source
	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// reading teh response body
	b, err := ioutil.ReadAll(resp.Body)

	// handling download in a different goroutine
	go func(fp string, b []byte) {
		// creating directories for cache document
		err := os.MkdirAll(path.Join(".", "cache", id), os.ModePerm)
		if err != nil {
			logrus.Errorf("couldn't create directory, %v\n", err)
		}

		// writing file that contains cache document
		err = ioutil.WriteFile(fp, b, os.ModePerm)
		if err != nil {
			logrus.Errorf("couldn't write file, %v\n", err)
		}
	}(fp, b)

	// returning the data and error
	return b, err
}
