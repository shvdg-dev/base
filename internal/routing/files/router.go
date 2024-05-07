package files

import (
	"base/internal/handlers/files"
	"net/http"
)

// FileServer sets up a file server for serving static files.
func FileServer() {
	http.Handle("/public/", files.Handler())
}
