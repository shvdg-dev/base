package routes

import "base/internal/routes/home"

func SetupRouter() {
	FileServer()
	home.Router()
}
