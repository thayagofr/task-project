package request

import "task-api/internal/core/domain"

type LoginCredentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (credentials *LoginCredentials) ToDomain() *domain.UserCredentials {
	return &domain.UserCredentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	}
}
