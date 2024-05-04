package login

import (
	"base/internal/handlers/login"
	"net/http"
)

func Router() {
	http.HandleFunc("/login", login.Navigation)
}
