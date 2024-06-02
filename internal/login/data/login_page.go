package data

// LoginPage holds all the data relevant for the LoginPage page.
type LoginPage struct {
	Email    string
	Password string
}

// LoginOption is a function that applies a specific configuration to a LoginPage instance.
type LoginOption func(*LoginPage)

// WithEmail is the option setter for the email field.
func WithEmail(email string) LoginOption {
	return func(d *LoginPage) {
		d.Email = email
	}
}

// WithPassword is the option setter for the password field.
func WithPassword(password string) LoginOption {
	return func(d *LoginPage) {
		d.Password = password
	}
}

// NewLogin creates a new instance of LoginPage with some Options.
func NewLogin(options ...LoginOption) *LoginPage {
	data := &LoginPage{}
	for _, option := range options {
		option(data)
	}
	return data
}
