package inmemory

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"task-api/internal/adapters/out/database/inmemory/entities"
	"task-api/internal/adapters/out/database/mongodb/documents"
	"task-api/internal/core/domain"
	"task-api/internal/ports/out"
	"time"
)

var _ out.UserProvider = &UserInMemDBRepository{}

type UserInMemDBRepository struct {
	usersDB *sync.Map
}

func NewUserInMemDBRepository() *UserInMemDBRepository {
	return &UserInMemDBRepository{usersDB: new(sync.Map)}
}

func (repo *UserInMemDBRepository) FindCredentialsByEmail(_ context.Context, email string) (*domain.UserCredentials, error) {
	var (
		exists      bool
		credentials domain.UserCredentials
	)

	repo.usersDB.Range(func(key, value any) bool {
		user := value.(*entities.User)
		if user.Email == email {
			exists = true
			credentials = domain.UserCredentials{
				Email:    user.Email,
				Password: user.Password,
			}
			return false
		}
		// Continue iterating
		return true
	})

	if exists {
		return &credentials, nil
	}

	return nil, out.ErrUserNotFound
}

func (repo *UserInMemDBRepository) RegisterUser(_ context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error) {
	user := entities.FromDomain(newUser)
	user.ID = uuid.NewString()
	user.CreationDate = time.Now()

	repo.usersDB.Store(user.ID, user)

	return &domain.RegisteredUser{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (repo *UserInMemDBRepository) ExistsByEmail(_ context.Context, email string) (bool, error) {
	var exists bool

	repo.usersDB.Range(func(key, value any) bool {
		if value.(*documents.User).Email == email {
			exists = true
			return false
		}
		// Continue iterating
		return true
	})

	if exists {
		return true, nil
	}

	return false, nil
}
