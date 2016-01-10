package main

import (
	"errors"
	"os"
)

var indexFile string

func init() {

	dir, _ := os.Getwd()
	newVal := "\"" + stripPrefix(dir+"/controllers/common") + "\""

	tmp := `package index

import (
	"net/http"

	` + newVal + `
)

// View blah blah
func View(w http.ResponseWriter, r *http.Request) {
	common.View(w, r, "index", struct{ Title string }{"Home"})
}
`

	indexFile = tmp

}

func appendIndex() error {
	file, err := os.Create("./controllers/index/index.go")
	if err != nil {
		return errors.New("WebGen: Failed to index controller: " + err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(indexFile))
	if err != nil {
		return errors.New("WebGen: Failed to index controller: " + err.Error())
	}

	return nil
}
