package home

import (
	"base/internal/handlers/home"
	"net/http"
)

func Router() {
	http.HandleFunc("/", home.Navigation)
	http.HandleFunc("/home", home.Navigation)
}
