package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if err := appendFolders(); err != nil {
		log.Println(err)
	}
	if err := appendIndex(); err != nil {
		log.Println(err)
	}
	if err := appendMain(); err != nil {
		log.Println(err)
	}
	if err := appendViews(); err != nil {
		log.Println(err)
	}
	if err := appendCommon(); err != nil {
		log.Println(err)
	}
	log.Println("Created!")
}

func stripPrefix(val string) string {
	if yes := strings.Contains(val, os.Getenv("GOPATH")); yes {
		newPath := strings.Replace(val, os.Getenv("GOPATH")+"/src/", "", 3)
		return newPath
	}

	log.Println("You have either not set your Gopath env variable or you are not generating inside your 'githhub.com/<your acc>/<folder>'.. The import for common in `./main.go` and `./controllers/index/index.go` might be messed up, please re-import the `./controllers/common/common.go` file to those two files for this to work properly.")
	return ""
}
