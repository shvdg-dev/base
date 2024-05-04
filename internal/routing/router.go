package routing

import (
	"base/internal/routing/docs"
	"base/internal/routing/home"
	"base/internal/routing/login"
)

func SetupRouter() {
	FileServer()
	home.Router()
	docs.Router()
	login.Router()
}
