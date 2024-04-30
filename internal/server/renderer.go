package server

import (
	"base/internal/views/document"
	. "github.com/maragudk/gomponents"
	"net/http"
)

// render renders a component, and is wrapped in a document when no swapping is intended.
func render(title string, content Node, w http.ResponseWriter, r *http.Request) {
	target := r.Header.Get("HX-Target")

	var err error
	if target != "" {
		err = content.Render(w)
	} else {
		err = document.Document(title, content).Render(w)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
