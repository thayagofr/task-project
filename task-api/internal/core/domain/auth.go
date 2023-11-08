package domain

import "time"

type UserCredentials struct {
	Email    string
	Password string
}

type AccessCredentials struct {
	Token          string
	Type           string
	ExpirationDate time.Time
}
