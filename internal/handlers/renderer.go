package handlers

import (
	"base/internal/models"
	"base/internal/views/document"
	. "github.com/maragudk/gomponents"
	"net/http"
)

// Render renders a component, and is wrapped in a document when no swapping is intended.
func Render(page models.Page, content Node, writer http.ResponseWriter, request *http.Request) {
	target := request.Header.Get("HX-Target")

	var err error
	if target != "" {
		writer.Header().Set("H-Title", page.Title)
		err = document.Slot(content).Render(writer)
	} else {
		err = document.Document(page, content).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
