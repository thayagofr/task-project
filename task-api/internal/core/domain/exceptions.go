package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailAlreadyUsed   = errors.New("email already used")
	ErrEmailRequired      = errors.New("email is required")
	ErrTaskNameRequired   = errors.New("task name is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrPasswordTooShort   = errors.New("password is too short")
)
