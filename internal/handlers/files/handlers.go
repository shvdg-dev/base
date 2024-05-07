package files

import "net/http"

func Handler() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}
