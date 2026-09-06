package tasks_service

import (
	"context"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	taskID int,
	patch core_domain.TaskPatch,
) (core_domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return core_domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	patchedTask, err := s.tasksRepository.PatchTask(ctx, taskID, task)
	if err != nil {
		return core_domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}
