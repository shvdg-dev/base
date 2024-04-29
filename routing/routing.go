package routing

import (
	"net/http"
	"tab-collector/routing/handlers"
)

func SetupRouting() {
	http.HandleFunc("/playlist", handlers.Playlist)
}
