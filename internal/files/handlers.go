package files

import "net/http"

// Handler handles HTTP requests to serve static files from a directory
func Handler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
}
