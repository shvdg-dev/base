package handlers

import (
	"base/internal/views/document"
	. "github.com/maragudk/gomponents"
	"net/http"
)

// Render renders a component, and is wrapped in a document when no swapping is intended.
func Render(title string, content Node, writer http.ResponseWriter, request *http.Request) {
	target := request.Header.Get("HX-Target")

	var err error
	if target != "" {
		err = document.Slot(content).Render(writer)
	} else {
		err = document.Document(title, content).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
