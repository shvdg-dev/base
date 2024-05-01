package routes

import "net/http"

// FileServer sets up a file server for serving static files.
func FileServer() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
