package domain

import (
	"time"
)

type UserToRegister struct {
	Age      *int
	Name     string
	Email    string
	Password string
}

type RegisteredUser struct {
	ID             string
	Name           string
	Age            *int
	Email          string
	RegisteredDate time.Time
}

type User struct {
	ID             string
	Name           string
	Age            *int
	Email          string
	Password       string
	RegisteredDate time.Time
}
