package raw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/ritwik310/my-website/server/db"
)

// GetDocument .
func GetDocument(id string, src string, doctype int8) ([]byte, error) {
	// constructing cahce filename and getting the actual source file url
	var fn string
	switch doctype {
	case db.DocTypeMD:
		fn = id + ".md"
	case db.DocTypeHTML:
		fn = id + ".html"
	default:
		return nil, fmt.Errorf("document type not defined")
	}

	// checking if requested file Exists or Not
	if _, err := os.Stat(path.Join(".", "cache", fn)); err == nil {
		return ioutil.ReadFile(fn)
	}

	// reading data from remote source
	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
