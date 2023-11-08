package usecase

import (
	"context"
	"task-api/internal/core/domain"
	"task-api/internal/ports/out"
)

type TasksManager struct {
	userProvider   out.UserProvider
	taskRepository out.TaskRepository
}

func (manager *TasksManager) Register(ctx context.Context, taskToRegister *domain.TaskToRegister) (*domain.RegisteredTask, error) {
	if taskToRegister.Name == "" {
		return nil, domain.ErrTaskNameRequired
	}

	ownerExists, err := manager.userProvider.ExistsByID(ctx, taskToRegister.OwnerID)
	if err != nil {
		return nil, err
	}

	if !ownerExists {
		return nil, domain.ErrInvalidCredentials
	}

	taskToRegister.

}
