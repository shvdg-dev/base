package info

import (
	"base/internal/sessions"
	"net/http"
)

// Informer provides information about a particular request.
type Informer struct {
	Sessions *sessions.Service
}

// NewInformant creates a new Informant instance.
func NewInformant(sessions *sessions.Service) *Informer {
	return &Informer{Sessions: sessions}
}

// NewInfo creates a new Info object based on the provided request and options.
func (i *Informer) NewInfo(request *http.Request, opts ...Option) *Info {
	info := &Info{Path: request.URL.Path, IsAuthenticated: i.Sessions.IsAuthenticated(request)}
	for _, opt := range opts {
		opt(info)
	}
	return info
}
