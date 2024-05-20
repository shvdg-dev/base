package docs

import (
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"net/http"
)

func (d *Docs) HandleDocsPage(writer http.ResponseWriter, request *http.Request) {
	rend.GetRenderer().Render(
		doc.NewInfo(request, doc.WithTitle("Docs")), d.Page(),
		writer, request)
}
