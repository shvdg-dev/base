package renderer

import (
	consts "base/internal/constants"
	ctx "base/internal/context"
	doc "base/internal/document"
	"base/internal/document/info"
	vi "base/internal/views"
	. "github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"net/http"
)

// Renderer is responsible for rendering documents and partials.
type Renderer struct {
	Document *doc.Document
}

// NewRenderer returns a new instance of Renderer.
func NewRenderer(context *ctx.Context, views *vi.Views) *Renderer {
	return &Renderer{Document: doc.NewDocument(context, views)}
}

// Render renders the content to the HTTP response writer based on the provided information and content nodes.
func (r *Renderer) Render(writer http.ResponseWriter, request *http.Request, info *info.Info, content ...Node) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set(consts.HeaderTitle, info.Title)
		writer.Header().Set(consts.HeaderPath, info.Path)
		err = r.Document.CreatePartial(r.Document.Views.Navbar.CreateNavItems(info), Group(content)).Render(writer)
	} else {
		err = r.Document.CreateDocument(info, Group(content)).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
