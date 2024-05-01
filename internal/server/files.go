package server

import "net/http"

// fileServer sets up a file server for serving static files.
func fileServer() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
