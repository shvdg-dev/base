package document

import "net/http"

type InfoOption func(*Info)

type Info struct {
	Path   string
	Title  string
	Errors []string // List of error messages
}

func WithTitle(t string) InfoOption {
	return func(i *Info) {
		i.Title = t
	}
}

func WithErrors(e []string) InfoOption {
	return func(i *Info) {
		i.Errors = e
	}
}

func NewInfo(request *http.Request, opts ...InfoOption) *Info {
	info := &Info{Path: request.URL.Path}
	for _, opt := range opts {
		opt(info)
	}
	return info
}

func (i *Info) HasErrors() bool {
	return len(i.Errors) > 0
}
