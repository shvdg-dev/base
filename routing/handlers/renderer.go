package handlers

import (
	"net/http"
	"tab-collector/models"
	"tab-collector/views/layout"
)

// Render renders a (partial) page.
func Render(page models.Page, w http.ResponseWriter, r *http.Request) {
	target := r.Header.Get(TargetHeader)

	var err error
	if target != "" {
		err = layout.Partial(page.Content).Render(w)
	} else {
		err = layout.Page(page).Render(w)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
