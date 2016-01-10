package main

import (
	"errors"
	"os"
)

var mainFile string

func init() {
	dir, _ := os.Getwd()
	commonVal := "\"" + stripPrefix(dir+"/controllers/common") + "\""
	indexVal := "\"" + stripPrefix(dir+"/controllers/index") + "\""

	tmp := `package main

import (
	"log"
	"net/http"

    ` + commonVal + `
    ` + indexVal + `
	"github.com/gorilla/mux"
)

func main() {

	// Declare New Router
	listenPort := ":8080"
	router := mux.NewRouter()

	// Create Middleware to log any sort of requests
	http.Handle("/", common.Log(router))
	// Serve Static Files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route Controllers
	router.HandleFunc("/", index.View).Methods("GET")

	// Serve And Listen On Port
	log.Println("Listening And Serving On Port", listenPort)
	http.ListenAndServe(listenPort, nil)

}`

	mainFile = tmp
}

func appendMain() error {
	file, err := os.Create("./main.go")
	if err != nil {
		return errors.New("WebGen: Failed to create index view" + err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(mainFile))
	if err != nil {
		return errors.New("WebGen: Failed to create index view" + err.Error())
	}
	file.Close()

	return nil
}
