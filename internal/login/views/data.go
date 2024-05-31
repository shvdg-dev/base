package views

// Data holds all the data relevant for the Login page.
type Data struct {
	CurrentEmail    string
	CurrentPassword string
}

// NewData creates a new instance of Data.
func (l *Login) NewData(currentEmail, currentPassword string) *Data {
	return &Data{CurrentEmail: currentEmail, CurrentPassword: currentPassword}
}
