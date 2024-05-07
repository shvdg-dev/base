package websocket

import (
	"base/internal/handlers/websocket"
	"net/http"
)

func Router() {
	http.HandleFunc("/websocket", websocket.Handler)
}
