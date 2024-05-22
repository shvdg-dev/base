package renderer

import (
	ctx "base/internal/context"
	doc "base/internal/document"
	"base/internal/document/info"
	vi "base/internal/views"
	. "github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"net/http"
)

type Renderer struct {
	Document *doc.Document
}

func NewRenderer(context *ctx.Context, views *vi.Views) *Renderer {
	return &Renderer{Document: doc.NewDocument(context, views)}
}

// Render renders the provided components normally or in a new Document when no target is defined.
func (r *Renderer) Render(writer http.ResponseWriter, request *http.Request, info *info.Info, content ...Node) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set("H-Title", info.Title)
		writer.Header().Set("H-Path", info.Path)
		err = r.Document.CreatePartial(r.Document.Views.Navbar.CreateNavItems(info), Group(content)).Render(writer)
	} else {
		err = r.Document.CreateDocument(request, info, Group(content)).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
