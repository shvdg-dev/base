package docs

import (
	doc "base/internal/document"
	"net/http"
)

func (d *Docs) handleDocsPage(writer http.ResponseWriter, request *http.Request) {
	d.context.Renderer.Render(doc.NewInfo(doc.WithPath("/docs"), doc.WithTitle("Docs")), d.Page(), writer, request)
}
