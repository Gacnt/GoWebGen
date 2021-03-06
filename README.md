# Go Web Gen
Go Web Gen creates a skeleton for a go app, following this structure:

```
.
├── controllers
│   ├── common
│   │   └── common.go
│   └── index
│       └── index.go
├── main.go
├── static
│   ├── css
│   ├── img
│   └── js
└── views
    ├── _footer.html
    ├── _header.html
    └── index
        └── index.html
```

#Prerequisites
**GoWebGen** assumes you have set your `$GOPATH`, and are running the command from within your `$GOPATH/src/github.com/<Your Account>/<Your New Project>`

#Installation
Just run
```
go get github.com/Gacnt/GoWebGen
```

#Usage
Create your new project folder within:
`$GOPATH/src/github.com/<Your Account>/<Your New Project>`

When inside your new project simply run:

```
GoWebGen
```

You will see 

>"Created!"

Navigate to "/controllers/common/common.go" and go to: 
```
var store = sessions.NewCookieStore([]byte("your-key"))`
```

replace `your-key` with something more secure.

Then you will be good to go!

# Extras
Includes helpers like:
```
common.View(w http.ResponseWriter, r *http.Request, tmplN string, data interface{}) // Renders a template
common.SendJSON(w http.ResponseWriter, data interface{}) // Sends a JSON response
common.Sesh(r *http.Request) *sessions.Session // Retrieves the gorilla session
```

# FAQ
- How do I create a new template file
  - Simply create your file, and follow [this](#tstruct) structure
- How do I render include static css/js/img files
  - In your template, just include as you normally would with the path `/static/<js|img|css>/yourFile.<js|img|css>


# Template Structure <a name='tstruct'/>

In the `define` tag put whatever you want for a name, you can later render this template from a controller by calling:

```
common.View(w, r, "name", nil)
```

```
{ define "newName" }} {{ template "_header" . }}
    // Put whatever you want inside here 
{{ template "_footer" . }} {{ end }}
```
