package renderer

import (
	doc "base/internal/document"
	"base/pkg/i18n"
	"github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"net/http"
)

type Renderer struct {
	Localizer *i18n.Localizer
	Document  *doc.Document
}

func NewRenderer(localizer *i18n.Localizer) *Renderer {
	return &Renderer{Localizer: localizer, Document: doc.NewDocument(localizer)}
}

// Render renders a component in either the content slot or in a new Document when no target is defined.
func (r *Renderer) Render(info *doc.Info, content gomponents.Node, writer http.ResponseWriter, request *http.Request) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set("H-Title", info.Title)
		err = r.Document.CreatePartial(content).Render(writer)
	} else {
		err = r.Document.CreateDocument(info, content).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
