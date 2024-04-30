package server

import (
	"base/internal/models"
	"base/internal/views/document"
	"net/http"
)

// render renders a component, and is wrapped in a document when no swapping is intended.
func render(component models.Component, w http.ResponseWriter, r *http.Request) {
	target := r.Header.Get("HX-Target")

	var err error
	if target != "" {
		err = component.Content.Render(w)
	} else {
		err = document.Document(component).Render(w)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
