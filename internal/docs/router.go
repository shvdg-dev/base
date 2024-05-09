package docs

import (
	"net/http"
)

func Router() {
	http.HandleFunc("/docs", Handler)
}
