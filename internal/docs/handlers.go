package docs

import (
	consts "base/internal/constants"
	doc "base/internal/document/info"
	"net/http"
)

// HandleDocsPage handles the documentation page request.
func (d *Docs) HandleDocsPage(writer http.ResponseWriter, request *http.Request) {
	info := d.Context.Informer.NewInfo(request, doc.WithTitle(d.Context.Localizer.Localize(consts.BundleDocs)))
	page := d.Views.Docs.CreateDocsPage()
	d.Renderer.Render(writer, request, info, page)
}
