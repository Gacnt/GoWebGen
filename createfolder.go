package main

import (
	"log"
	"os"
)

func appendFolders() error {
	if err := os.MkdirAll("./views/index", 0777); err != nil {
		log.Println(err)
	}
	if err := os.MkdirAll("./controllers/common", 0777); err != nil {
		log.Println(err)
	}
	if err := os.Mkdir("./controllers/index", 0777); err != nil {
		log.Println(err)
	}
	if err := os.MkdirAll("./static/js", 0777); err != nil {
		log.Println(err)
	}
	if err := os.Mkdir("./static/css", 0777); err != nil {
		log.Println(err)
	}
	if err := os.Mkdir("./static/img", 0777); err != nil {
		log.Println(err)
	}

	return nil
}
