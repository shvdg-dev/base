package docs

import (
	doc "base/internal/document/info"
	"net/http"
)

func (d *Docs) HandleDocsPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Docs"))
	d.Renderer.Render(writer, request, info, d.Views.Docs.CreateDocsPage())
}
