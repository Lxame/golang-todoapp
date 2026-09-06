package tasks_service

import (
	"context"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	task core_domain.Task,
) (core_domain.Task, error) {
	if err := task.Validate(); err != nil {
		return core_domain.Task{}, fmt.Errorf(
			"validate task domain: %w",
			err,
		)
	}

	task, err := s.tasksRepository.CreateTask(ctx, task)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf(
			"create task: %w",
			err,
		)
	}

	return task, nil
}
