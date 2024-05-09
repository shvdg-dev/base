package login

import (
	"net/http"
)

func Router() {
	http.HandleFunc("/login", Handler)
}
