package home

import (
	"base/internal/handlers/home"
	"net/http"
)

func Router() {
	http.HandleFunc("/", home.Handler)
	http.HandleFunc("/home", home.Handler)
}
