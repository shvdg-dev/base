package handlers

import (
	"net/http"
	"tab-collector/models"
	"tab-collector/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Render(models.Page{
		Title:   "Home",
		Content: views.Home(),
	}, w, r)
}
