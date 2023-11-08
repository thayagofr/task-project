package out

import (
	"context"
	"task-api/internal/core/domain"
)

type TaskRepository interface {
	Register(ctx context.Context, taskToRegister *domain.TaskToRegister) (*domain.RegisteredTask, error)
}
