package files

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
}
