package document

type InfoOption func(*Info)

type Info struct {
	Path   string
	Title  string
	Errors []error
}

func WithPath(p string) InfoOption {
	return func(i *Info) {
		i.Path = p
	}
}

func WithTitle(t string) InfoOption {
	return func(p *Info) {
		p.Title = t
	}
}

func WithErrors(e []error) InfoOption {
	return func(p *Info) {
		p.Errors = e
	}
}

func NewInfo(opts ...InfoOption) *Info {
	page := &Info{}
	for _, opt := range opts {
		opt(page)
	}
	return page
}

func (p *Info) HasErrors() bool {
	return len(p.Errors) > 0
}
