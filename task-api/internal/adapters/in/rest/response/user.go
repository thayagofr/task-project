package response

import (
	"task-api/internal/core/domain"
	"time"
)

type RegisteredUser struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Age            *int      `json:"age"`
	Email          string    `json:"email"`
	RegisteredDate time.Time `json:"registered_date"`
}

func FromUserDomain(user *domain.RegisteredUser) *RegisteredUser {
	return &RegisteredUser{
		ID:             user.ID,
		Name:           user.Name,
		Age:            user.Age,
		Email:          user.Email,
		RegisteredDate: user.RegisteredDate,
	}
}
