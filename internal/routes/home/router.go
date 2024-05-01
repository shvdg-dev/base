package home

import (
	"net/http"
)

func Router() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/home", Handler)
}
