package out

import "task-api/internal/core/domain"

var _ TokenManager = &TokenManagerMock{}

type TokenManagerMock struct {
	MockedGenerate func(user *domain.User) (*domain.AccessCredentials, error)
	MockedValidate func(token string) error
}

func NewTokenProviderMock() *TokenManagerMock {
	return &TokenManagerMock{}
}

func (provider *TokenManagerMock) Generate(user *domain.User) (*domain.AccessCredentials, error) {
	if provider.MockedGenerate != nil {
		return provider.MockedGenerate(user)
	}
	return nil, nil
}

func (provider *TokenManagerMock) Validate(token string) error {
	if provider.MockedValidate != nil {
		return provider.MockedValidate(token)
	}
	return nil
}
