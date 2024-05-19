package users

import "time"

type User struct {
	ID             string
	Email          string
	Authorizations []Authorization
}

type Authorization struct {
	Name   string
	Expiry time.Time
}
