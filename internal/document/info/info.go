package info

import (
	"base/internal/sessions"
	"net/http"
)

type Informer struct {
	Sessions *sessions.Service
}

func NewInformant(sessions *sessions.Service) *Informer {
	return &Informer{Sessions: sessions}
}

type Option func(*Info)

type Info struct {
	Path            string
	Title           string
	IsAuthenticated bool
	Errors          []string // List of error messages
}

func WithTitle(t string) Option {
	return func(i *Info) {
		i.Title = t
	}
}

func WithErrors(e []string) Option {
	return func(i *Info) {
		i.Errors = e
	}
}

func (i *Informer) NewInfo(request *http.Request, opts ...Option) *Info {
	info := &Info{Path: request.URL.Path, IsAuthenticated: i.Sessions.IsAuthenticated(request)}
	for _, opt := range opts {
		opt(info)
	}
	return info
}

func (i *Info) HasErrors() bool {
	return len(i.Errors) > 0
}
