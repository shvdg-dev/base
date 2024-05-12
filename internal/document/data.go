package document

type PageOption func(*Page)

type Page struct {
	Path   string
	Title  string
	Errors []error
}

func WithTitle(t string) PageOption {
	return func(p *Page) {
		p.Title = t
	}
}

func WithErrors(e []error) PageOption {
	return func(p *Page) {
		p.Errors = e
	}
}

func NewPage(path string, opts ...PageOption) *Page {
	page := &Page{Path: path}
	for _, opt := range opts {
		opt(page)
	}
	return page
}
