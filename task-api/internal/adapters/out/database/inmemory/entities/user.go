package entities

import (
	"task-api/internal/core/domain"
	"time"
)

type User struct {
	ID           string
	Age          *int
	Name         string
	Email        string
	Password     string
	CreationDate time.Time
}

func FromDomain(user *domain.UserToRegister) *User {
	return &User{
		Age:          user.Age,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		CreationDate: time.Now(),
	}
}
