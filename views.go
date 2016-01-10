package main

import (
	"errors"
	"os"
)

var headerFile = `{{ define "_header" }}

<!DOCTYPE html>
<html>

<head>
    <title>
        {{ .Title }}
    </title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css">
    <link rel="stylesheet" href="https://cdn.rawgit.com/twbs/bootstrap/v4-dev/dist/css/bootstrap.css">
    <link rel="stylesheet" href="/static/css/custom.css">
</head>

<body>
{{ end }}
`

var footerFile = `{{ define "_footer" }}

<script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.1.1/js/tether.min.js"></script>
<script src="https://cdn.rawgit.com/twbs/bootstrap/v4-dev/dist/js/bootstrap.js"></script>
</body>

</html>

{{ end }}
`

var indexViewFile = `{{ define "index" }} {{ template "_header" . }}
<div class="container-fluid">
    <div class="row">
        <div class="col-md-10 col-centered text-center">
            <h3>Enjoy!</h3>
        </div>
    </div>
</div>

{{ template "_footer" . }} {{ end }}
`

func appendViews() error {
	file, err := os.Create("./views/_header.html")
	if err != nil {
		return errors.New("WebGen: Failed to create _header view" + err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(headerFile))
	if err != nil {
		return errors.New("WebGen: Failed to create _header view" + err.Error())
	}
	file.Close()

	file, err = os.Create("./views/_footer.html")
	if err != nil {
		return errors.New("WebGen: Failed to create _footer view" + err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(footerFile))
	if err != nil {
		return errors.New("WebGen: Failed to create _footer view" + err.Error())
	}
	file.Close()

	file, err = os.Create("./views/index/index.html")
	if err != nil {
		return errors.New("WebGen: Failed to create index view" + err.Error())
	}
	defer file.Close()

	_, err = file.Write([]byte(indexViewFile))
	if err != nil {
		return errors.New("WebGen: Failed to create index view" + err.Error())
	}
	file.Close()

	return nil
}
