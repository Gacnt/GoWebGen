package main

import (
	"errors"
	"log"
	"os"
)

var commonFile = `
package common

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-key"))
var templates = template.New("template").Funcs(template.FuncMap{
	"authenticated": func() bool {
		return false
	},
	"username": func() string {
		return ""
	},
})

type requestParams struct {
	ip, method, uri, protocol, host string
	elapsedTime                     time.Duration
}

func init() {
	filepath.Walk("views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			templates.ParseFiles(path)
		}

		return nil
	})
	log.Println("Templates Parsed")
}
func Log(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Time At Start Of Request
		StartTime := time.Now()
		// Serve Request
		router.ServeHTTP(w, r)
		// Get Time At End Of Request Handled
		EndTime := time.Now()
		// Get Time Difference Between Two Points
		elapsedTime := EndTime.Sub(StartTime)

		clientIP := r.RemoteAddr

		record := &requestParams{
			ip:          clientIP,
			method:      r.Method,
			uri:         r.RequestURI,
			protocol:    r.Proto,
			host:        r.Host,
			elapsedTime: elapsedTime,
		}

		logRecord := "" + record.ip + " " + record.protocol + " " + record.method + ": " + record.uri + ", host: " + record.host + " (load time: " + strconv.FormatFloat(record.elapsedTime.Seconds(), 'f', 5, 64) + " seconds)"

		log.Println(logRecord)

	})
}

func View(w http.ResponseWriter, r *http.Request, tmplN string, data interface{}) {

	err := templates.Funcs(template.FuncMap{
		"username": func() string {
			session := Sesh(r)
			username := ""
			if session.Values["username"] != nil {
				username = session.Values["username"].(string)
			}
			return username
		},
		"authenticated": func() bool {
			session := Sesh(r)
			authenticated := false
			if session.Values["authenticated"] != nil {
				authenticated = session.Values["authenticated"].(bool)
			}
			return authenticated
		},
	}).ExecuteTemplate(w, tmplN, data)

	if err != nil {
		log.Println("Error " + err.Error())
	}
}

// SendJSON Will send out a json message to the client
func SendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	msg, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Write(msg)
}

func Sesh(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "default-session")
	if err != nil {
		log.Printf("\nSESSION ERROR: %s", err.Error())
	}

	return session
}

`

func appendCommon() error {
	file, err := os.Create("./controllers/common/common.go")
	if err != nil {
		return errors.New("WebGen: Failed to common controller: " + err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(commonFile))
	if err != nil {
		log.Println(err)
	}

	return nil
}
