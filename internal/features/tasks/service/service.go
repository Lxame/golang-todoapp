package tasks_service

import (
	"context"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task core_domain.Task,
	) (core_domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]core_domain.Task, error)

	GetTask(
		ctx context.Context,
		taskID int,
	) (core_domain.Task, error)

	DeleteTask(
		ctx context.Context,
		taskID int,
	) error

	PatchTask(
		ctx context.Context,
		taskID int,
		task core_domain.Task,
	) (core_domain.Task, error)
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
