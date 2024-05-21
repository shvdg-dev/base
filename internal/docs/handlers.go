package docs

import (
	doc "base/internal/document/info"
	rend "base/internal/renderer"
	"net/http"
)

func (d *Docs) HandleDocsPage(writer http.ResponseWriter, request *http.Request) {
	info := doc.NewInfo(request, doc.WithTitle("Docs"))
	rend.GetRenderer().Render(writer, request, info, d.Views.Docs.CreateDocsPage())
}
