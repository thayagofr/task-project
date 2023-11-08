package usecase

import (
	"context"
	"errors"
	"task-api/internal/core/domain"
	"task-api/internal/ports/out"
)

type AuthUseCase struct {
	userProvider       out.UserProvider
	tokenProvider      out.TokenManager
	notificationSender out.NotificationSender
	passwordEncryptor  out.PasswordEncryptor
}

func NewAuthUseCase(
	userProvider out.UserProvider,
	tokenProvider out.TokenManager,
	notificationSender out.NotificationSender,
	passwordEncryptor out.PasswordEncryptor,
) *AuthUseCase {
	return &AuthUseCase{
		userProvider:       userProvider,
		tokenProvider:      tokenProvider,
		notificationSender: notificationSender,
		passwordEncryptor:  passwordEncryptor,
	}
}

func (useCase *AuthUseCase) Register(ctx context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error) {
	if err := useCase.validateCredentials(&domain.UserCredentials{
		Email:    newUser.Email,
		Password: newUser.Password,
	}); err != nil {
		return nil, err
	}

	exists, err := useCase.userProvider.ExistsByEmail(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrEmailAlreadyUsed
	}

	newUser.Password, err = useCase.passwordEncryptor.Encrypt(newUser.Password)
	if err != nil {
		return nil, err
	}

	registeredUser, err := useCase.userProvider.RegisterUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	if err = useCase.notificationSender.SendTo(ctx, domain.WelcomeToTaskIO, newUser.Email); err != nil {
		return nil, err
	}

	return registeredUser, nil
}

func (useCase *AuthUseCase) Authenticate(ctx context.Context, credentials *domain.UserCredentials) (*domain.AccessCredentials, error) {
	if err := useCase.validateCredentials(&domain.UserCredentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	}); err != nil {
		return nil, err
	}

	storedUser, err := useCase.userProvider.FindByEmail(ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, out.ErrUserNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
	}

	if !useCase.passwordEncryptor.Compare(storedUser.Password, credentials.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	return useCase.tokenProvider.Generate(storedUser)
}

func (useCase *AuthUseCase) validateCredentials(credentials *domain.UserCredentials) error {
	if credentials.Email == "" {
		return domain.ErrEmailRequired
	}

	if credentials.Password == "" {
		return domain.ErrPasswordRequired
	}

	if len(credentials.Password) < domain.MinPasswordLength {
		return domain.ErrPasswordTooShort
	}

	return nil
}
