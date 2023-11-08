package out

import (
	"errors"
	"task-api/internal/core/domain"
)

var (
	ErrGeneratingToken           = errors.New("error generating new user token")
	ErrInvalidTokenSigningMethod = errors.New("invalid token signing method")
	ErrInvalidToken              = errors.New("invalid token")
)

type TokenManager interface {
	Generate(user *domain.User) (*domain.AccessCredentials, error)
	Validate(token string) error
}
