package index

import (
	"net/http"

	"github.com/Gacnt/GoWebGen/_example/controllers/common"
)

// View blah blah
func View(w http.ResponseWriter, r *http.Request) {
	common.View(w, r, "index", struct{ Title string }{"Home"})
}
