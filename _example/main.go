package main

import (
	"log"
	"net/http"

    "github.com/Gacnt/GoWebGen/_example/controllers/common"
    "github.com/Gacnt/GoWebGen/_example/controllers/index"
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

}