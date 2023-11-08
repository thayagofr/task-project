package request

import "task-api/internal/core/domain"

type NewUser struct {
	Name     string `json:"name"`
	Age      *int   `json:"age"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *NewUser) ToDomain() *domain.UserToRegister {
	return &domain.UserToRegister{
		Age:      user.Age,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
