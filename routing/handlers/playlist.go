package handlers

import (
	"net/http"
	"tab-collector/models"
	"tab-collector/views"
)

func Playlist(w http.ResponseWriter, r *http.Request) {
	Render(models.Page{
		Title:   "Playlist",
		Content: views.Playlist(),
	}, w, r)
}
