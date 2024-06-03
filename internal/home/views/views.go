package views

import "base/internal/context"

// Home is used for views regarding the home page.
type Home struct {
	Context *context.Context
}

// NewHome initializes a new instance of Home.
func NewHome(context *context.Context) *Home {
	return &Home{Context: context}
}
