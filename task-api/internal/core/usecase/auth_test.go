package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"task-api/internal/core/domain"
	"task-api/internal/ports/out"
	"testing"
	"time"
)

func TestAuthUseCase_Register(t *testing.T) {

	var (
		ctx                      = context.Background()
		mockedUserProvider       = out.NewUUserProviderMock()
		mockedTokenProvider      = out.NewTokenProviderMock()
		mockedNotificationSender = out.NewNotificationSenderMock()
		mockedPasswordEncryptor  = out.NewPasswordEncryptorMock()
		useCase                  = NewAuthUseCase(
			mockedUserProvider,
			mockedTokenProvider,
			mockedNotificationSender,
			mockedPasswordEncryptor,
		)
	)

	t.Run("Case 1 - When the notification is empty then return ErrEmailRequired", func(t *testing.T) {
		newUser := &domain.UserToRegister{
			Email:    "",
			Password: "12345678",
		}
		_, got := useCase.Register(ctx, newUser)

		if !errors.Is(got, domain.ErrEmailRequired) {
			t.Failed()
		}
	})

	t.Run("Case 2 - When the password is empty then return ErrPasswordRequired", func(t *testing.T) {
		newUser := &domain.UserToRegister{
			Email:    "john@gmail.com",
			Password: "",
		}
		_, got := useCase.Register(ctx, newUser)

		if !errors.Is(got, domain.ErrPasswordRequired) {
			t.Failed()
		}
	})

	t.Run("Case 3 - When the password size is less than 8 then return ErrPasswordTooShort", func(t *testing.T) {
		newUser := &domain.UserToRegister{
			Email:    "john@gmail.com",
			Password: "12",
		}
		_, got := useCase.Register(ctx, newUser)

		if !errors.Is(got, domain.ErrPasswordTooShort) {
			t.Failed()
		}
	})

	t.Run("Case 4 - When the submitted notification is already in use than 8 then return ErrEmailAlreadyUsed", func(t *testing.T) {
		newUser := &domain.UserToRegister{
			Email:    "john@gmail.com",
			Password: "12345678",
		}

		mockedUserProvider.MockedExistsByEmail = func(ctx context.Context, email string) (bool, error) {
			return true, nil
		}
		defer func() {
			mockedUserProvider.MockedRegisterNewUser = nil
		}()

		_, got := useCase.Register(ctx, newUser)

		if !errors.Is(got, domain.ErrEmailAlreadyUsed) {
			t.Failed()
		}
	})

	t.Run("Case 5 - Success", func(t *testing.T) {
		newUser := &domain.UserToRegister{
			Name:     "John Doe",
			Email:    "john@gmail.com",
			Password: "12345678",
		}

		mockedUserProvider.MockedExistsByEmail = func(_ context.Context, _ string) (bool, error) {
			return false, nil
		}

		mockedPasswordEncryptor.MockedEncrypt = func(_ string) (string, error) {
			return "@@@@@@@@", nil
		}

		mockedUserProvider.MockedRegisterNewUser = func(_ context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error) {
			return &domain.RegisteredUser{
				ID:             uuid.NewString(),
				Name:           newUser.Name,
				Age:            newUser.Age,
				Email:          newUser.Email,
				RegisteredDate: time.Now(),
			}, nil
		}

		mockedNotificationSender.MockedSendTo = func(_ context.Context, _ string, _ string) error {
			return nil
		}

		defer func() {
			mockedUserProvider.MockedExistsByEmail = nil
			mockedUserProvider.MockedRegisterNewUser = nil
			mockedPasswordEncryptor.MockedEncrypt = nil
			mockedNotificationSender.MockedSendTo = nil
		}()

		_, got := useCase.Register(ctx, newUser)

		if got != nil {
			t.Failed()
		}
	})
}

func TestAuthUseCase_Authenticate(t *testing.T) {

	var (
		ctx                      = context.Background()
		mockedUserProvider       = out.NewUUserProviderMock()
		mockedTokenProvider      = out.NewTokenProviderMock()
		mockedNotificationSender = out.NewNotificationSenderMock()
		mockedPasswordEncryptor  = out.NewPasswordEncryptorMock()
		useCase                  = NewAuthUseCase(
			mockedUserProvider,
			mockedTokenProvider,
			mockedNotificationSender,
			mockedPasswordEncryptor,
		)
	)

	t.Run("Case 1 - When the notification is empty then return ErrEmailRequired", func(t *testing.T) {
		credentials := &domain.UserCredentials{
			Email:    "",
			Password: "12345678",
		}
		_, got := useCase.Authenticate(ctx, credentials)

		if !errors.Is(got, domain.ErrEmailRequired) {
			t.Failed()
		}
	})

	t.Run("Case 2 - When the password is empty then return ErrPasswordRequired", func(t *testing.T) {
		credentials := &domain.UserCredentials{
			Email:    "john@gmail.com",
			Password: "",
		}
		_, got := useCase.Authenticate(ctx, credentials)

		if !errors.Is(got, domain.ErrPasswordRequired) {
			t.Failed()
		}
	})

	t.Run("Case 3 - When the password size is less than 8 then return ErrPasswordTooShort", func(t *testing.T) {
		credentials := &domain.UserCredentials{
			Email:    "john@gmail.com",
			Password: "12",
		}
		_, got := useCase.Authenticate(ctx, credentials)

		if !errors.Is(got, domain.ErrPasswordTooShort) {
			t.Failed()
		}
	})

	t.Run("Case 4 - When there is not user with the specified email then return ErrInvalidCredentials", func(t *testing.T) {
		credentials := &domain.UserCredentials{
			Email:    "john@gmail.com",
			Password: "12345678",
		}

		mockedUserProvider.MockedFindByEmail = func(_ context.Context, _ string) (*domain.User, error) {
			return nil, out.ErrUserNotFound
		}

		defer func() {
			mockedUserProvider.MockedFindByEmail = nil
		}()

		_, got := useCase.Authenticate(ctx, credentials)

		if !errors.Is(got, domain.ErrInvalidCredentials) {
			t.Failed()
		}
	})

	t.Run("Case 5 - When the password doesnt match then return ErrInvalidCredentials", func(t *testing.T) {
		var (
			email       = "john@gmail.com"
			credentials = &domain.UserCredentials{
				Email:    email,
				Password: "PASSWORD",
			}
			storedUser = &domain.User{
				ID:       uuid.NewString(),
				Email:    email,
				Password: "HASHED_PASSWORD",
			}
		)

		mockedUserProvider.MockedFindByEmail = func(_ context.Context, _ string) (*domain.User, error) {
			return storedUser, nil
		}

		mockedPasswordEncryptor.MockedCompare = func(encrypted string, original string) bool {
			return false
		}

		defer func() {
			mockedUserProvider.MockedFindByEmail = nil
			mockedPasswordEncryptor.MockedCompare = nil
		}()

		_, got := useCase.Authenticate(ctx, credentials)

		if !errors.Is(got, domain.ErrInvalidCredentials) {
			t.Failed()
		}
	})

}
