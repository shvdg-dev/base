package routing

import (
	"base/internal/routing/handlers"
	"net/http"
)

func home() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/home", handlers.Home)
}
