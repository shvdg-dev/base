package renderer

import (
	ctx "base/internal/context"
	doc "base/internal/document"
	"base/internal/document/info"
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

func NewRenderer(context *ctx.Context) *Renderer {
	once.Do(func() {
		renderer = &Renderer{Document: doc.NewDocument(context)}
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
func (r *Renderer) Render(info *info.Info, content Node, writer http.ResponseWriter, request *http.Request) {
	target := hxhttp.GetTarget(request.Header)

	var err error
	if target != "" {
		writer.Header().Set("H-Title", info.Title)
		navBar := r.Document.NewNavBar(info, request)
		err = r.Document.CreatePartial(content, navBar.CreateMenuItems(), navBar.CreateInOutButton()).Render(writer)
	} else {
		err = r.Document.CreateDocument(info, content, request).Render(writer)
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
