package tasks_transport_http

import (
	"context"
	"net/http"

	core_domain "github.com/Lxame/golang-todoapp/internal/core/domain"
	core_http_server "github.com/Lxame/golang-todoapp/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
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
		patch core_domain.TaskPatch,
	) (core_domain.Task, error)
}

func NewTasksHTTPHandler(
	taskService TasksService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: taskService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: h.PatchTask,
		},
	}
}
