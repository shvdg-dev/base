package routing

import (
	"base/internal/routing/docs"
	"base/internal/routing/home"
)

func SetupRouter() {
	FileServer()
	home.Router()
	docs.Router()
}