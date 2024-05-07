package routing

import (
	"base/internal/routing/docs"
	"base/internal/routing/files"
	"base/internal/routing/home"
	"base/internal/routing/login"
	"base/internal/routing/websocket"
)

func SetupRouter() {
	files.FileServer()
	websocket.Router()
	home.Router()
	docs.Router()
	login.Router()
}
