package info

// Option can be used to modify the fields of an `Info` struct.
type Option func(*Info)

// Info represents information about a specific request or operation.
type Info struct {
	Path            string
	Title           string
	IsAuthenticated bool
	Errors          []string // List of error messages
}

// WithPath allows to set the path of a given Info struct.
func WithPath(path string) Option {
	return func(i *Info) {
		i.Path = path
	}
}

// WithTitle allows to set the title of an Info struct.
func WithTitle(title string) Option {
	return func(i *Info) {
		i.Title = title
	}
}

// WithErrors allows to set the errors field of the Info struct
func WithErrors(errors []string) Option {
	return func(i *Info) {
		i.Errors = errors
	}
}

// HasErrors returns true if there are any errors in the Info struct, otherwise false.
func (i *Info) HasErrors() bool {
	return len(i.Errors) > 0
}
