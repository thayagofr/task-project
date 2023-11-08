package in

import (
	"context"
	"task-api/internal/core/domain"
)

type RegisterUser interface {
	Register(ctx context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error)
	Authenticate(ctx context.Context, credentials *domain.UserCredentials) (*domain.AccessCredentials, error)
}
