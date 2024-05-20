package renderer

import (
	doc "base/internal/document"
	"base/internal/document/info"
	"base/internal/sessions"
	"base/pkg/i18n"
	"github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"net/http"
)

type Renderer struct {
	Document *doc.Document
}

func NewRenderer(localizer *i18n.Localizer, sessions *sessions.Service) *Renderer {
	return &Renderer{Document: doc.NewDocument(localizer, sessions)}
}

// Render renders a component in either the content slot or in a new Document when no target is defined.
func (r *Renderer) Render(info *info.Info, content gomponents.Node, writer http.ResponseWriter, request *http.Request) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set("H-Title", info.Title)
		err = r.Document.CreatePartial(content).Render(writer)
	} else {
		err = r.Document.CreateDocument(info, content, request).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
