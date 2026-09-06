package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
	core_errors "github.com/Lxame/golang-todoapp/internal/core/errors"
	core_postgres_pool "github.com/Lxame/golang-todoapp/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	taskID int,
) (core_domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT
			id,
			version,
			title,
			description,
			completed,
			created_at,
			completed_at,
			author_user_id
		FROM todoapp.tasks
		WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, taskID)
	var taskModel TaskModel
	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return core_domain.Task{}, fmt.Errorf(
				"task with id=`%d`: %w",
				taskID,
				core_errors.ErrNotFound,
			)
		}

		return core_domain.Task{}, fmt.Errorf("scan error: %w", err)
	}

	return taskDomainFromModel(taskModel), nil
}
