package routing

import (
	"base/models"
	"base/views/layout"
	"net/http"
)

// render renders a (partial) page.
func render(page models.Page, w http.ResponseWriter, r *http.Request) {
	target := r.Header.Get("HX-Target")

	var err error
	if target != "" {
		err = layout.Fragment(page.Content).Render(w)
	} else {
		err = layout.Document(page).Render(w)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
