package docs

import (
	doc "base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func (d *Docs) handleDocsPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(doc.NewInfo(doc.WithPath("/docs"), doc.WithTitle("Docs")), d.Page(), writer, request)
}
