package docs

import (
	doc "base/internal/document/info"
	"net/http"
)

func (d *Docs) HandleDocsPage(writer http.ResponseWriter, request *http.Request) {
	d.Context.Renderer.Render(
		doc.NewInfo(request, doc.WithTitle("Docs")), d.Page(),
		writer, request)
}
