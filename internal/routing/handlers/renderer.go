package handlers

import (
	"base/internal/views/document"
	. "github.com/maragudk/gomponents"
	"net/http"
)

// render renders a component, and is wrapped in a document when no swapping is intended.
func render(title string, content Node, writer http.ResponseWriter, request *http.Request) {
	target := request.Header.Get("HX-Target")

	var err error
	if target != "" {
		err = content.Render(writer)
	} else {
		err = document.Document(title, content).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
