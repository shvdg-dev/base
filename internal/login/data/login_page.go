package data

// LoginData holds all the data relevant for any login view.
type LoginData struct {
	Email    string
	Password string
}

// LoginOption is a function that applies a specific configuration to a LoginData instance.
type LoginOption func(*LoginData)

// WithEmail is the option setter for the email field.
func WithEmail(email string) LoginOption {
	return func(d *LoginData) {
		d.Email = email
	}
}

// WithPassword is the option setter for the password field.
func WithPassword(password string) LoginOption {
	return func(d *LoginData) {
		d.Password = password
	}
}

// NewLoginData creates a new instance of LoginData with some Options.
func NewLoginData(options ...LoginOption) *LoginData {
	data := &LoginData{}
	for _, option := range options {
		option(data)
	}
	return data
}
