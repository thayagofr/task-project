package out

import (
	"context"
	"errors"
	"task-api/internal/core/domain"
)

var (
	ErrUserNotFound = errors.New("user not found. invalid credentials")
)

type UserProvider interface {
	RegisterUser(ctx context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByID(ctx context.Context, ID string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
