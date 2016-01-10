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
**GoWebGen** assumes you have set your $GOPATH, and are running the command from with your `$GOPATH/src/github.com/<Your Account>/<Your New Project>`

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
webgen
```

You will see 

>"Created!"

Then you will be good to go!

