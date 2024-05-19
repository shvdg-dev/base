package renderer

import (
	"base/internal/document"
	"github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"net/http"
)

// Render renders a component in either the content slot or in a new document when no target is defined.
func Render(info *document.Info, content gomponents.Node, writer http.ResponseWriter, request *http.Request) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set("H-Title", info.Title)
		err = document.Slot(content).Render(writer)
	} else {
		err = document.Document(info, content).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
