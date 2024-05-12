package docs

import (
	"base/internal/document"
	"base/internal/renderer"
	"net/http"
)

func handleDocsPage(writer http.ResponseWriter, request *http.Request) {
	renderer.Render(document.NewPage("/docs", document.WithTitle("Docs")), page(), writer, request)
}
