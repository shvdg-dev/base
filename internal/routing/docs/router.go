package docs

import (
	"base/internal/handlers/docs"
	"net/http"
)

func Router() {
	http.HandleFunc("/docs", docs.Navigation)
}
