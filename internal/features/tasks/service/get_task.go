package tasks_service

import (
	"context"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

func (s *TasksService) GetTask(
	ctx context.Context,
	taskID int,
) (core_domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf(
			"get task from repository: %w",
			err,
		)
	}

	return task, nil
}
