package renderer

import (
	ctx "base/internal/context"
	doc "base/internal/document"
	"base/internal/document/info"
	vi "base/internal/views"
	. "github.com/maragudk/gomponents"
	hxhttp "github.com/maragudk/gomponents-htmx/http"
	"log"
	"net/http"
	"sync"
)

type Renderer struct {
	Document *doc.Document
}

var renderer *Renderer
var once sync.Once

func NewRenderer(context *ctx.Context, views *vi.Views) *Renderer {
	once.Do(func() {
		renderer = &Renderer{Document: doc.NewDocument(context, views)}
	})
	return renderer
}

func GetRenderer() *Renderer {
	if renderer == nil {
		log.Fatal("Renderer is not instantiated yet.")
	}
	return renderer
}

// Render renders a component in either the content slot or in a new Document when no target is defined.
func (r *Renderer) Render(writer http.ResponseWriter, request *http.Request, info *info.Info, content ...Node) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set("H-Title", info.Title)
		err = r.Document.CreatePartial(r.Document.Views.Navbar.CreateNavItems(info), Group(content)).Render(writer)
	} else {
		err = r.Document.CreateDocument(request, info, Group(content)).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
