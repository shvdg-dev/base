package routing

import (
	"net/http"
	"tab-collector/routing/handlers"
)

func SetupRouting() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/home", handlers.Home)
}
