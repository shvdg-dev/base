package files

import (
	"net/http"
)

// Router sets up a file server for serving static files.
func Router() {
	http.Handle("/public/", Handler())
}
