package out

import (
	"context"
	"task-api/internal/core/domain"
)

var _ UserProvider = &UserProviderMock{}

type UserProviderMock struct {
	MockedRegisterNewUser func(ctx context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error)
	MockedExistsByEmail   func(ctx context.Context, email string) (bool, error)
	MockedExistsByID      func(ctx context.Context, ID string) (bool, error)
	MockedFindByEmail     func(ctx context.Context, email string) (*domain.User, error)
}

func NewUUserProviderMock() *UserProviderMock {
	return &UserProviderMock{}
}

func (mock *UserProviderMock) ExistsByID(ctx context.Context, ID string) (bool, error) {
	if mock.MockedExistsByID != nil {
		return mock.MockedExistsByID(ctx, ID)
	}
	return false, nil
}

func (mock *UserProviderMock) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if mock.MockedFindByEmail != nil {
		return mock.MockedFindByEmail(ctx, email)
	}
	return nil, nil
}

func (mock *UserProviderMock) RegisterUser(ctx context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error) {
	if mock.MockedRegisterNewUser != nil {
		return mock.MockedRegisterNewUser(ctx, newUser)
	}
	return nil, nil
}

func (mock *UserProviderMock) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	if mock.MockedExistsByEmail != nil {
		return mock.MockedExistsByEmail(ctx, email)
	}
	return false, nil
}
