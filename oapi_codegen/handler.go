package oapicodegen

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	r TaskRepository
}

func NewTaskHandler(r TaskRepository) *TaskHandler {
	return &TaskHandler{r: r}
}

// CreateTask implements generated.ServerInterface.
func (t *TaskHandler) CreateTask(ctx echo.Context) error {
	panic("unimplemented")
}

// DeleteTask implements generated.ServerInterface.
func (t *TaskHandler) DeleteTask(ctx echo.Context, id int) error {
	panic("unimplemented")
}

// GetTasks implements generated.ServerInterface.
func (t *TaskHandler) GetTasks(ctx echo.Context) error {
	res, err := t.r.FetchTasks()
	if err != nil {
		return err
	}
	slog.Info("Result", "data", res)
	return ctx.JSON(http.StatusOK, res)
}

// PartiallyUpdateTask implements generated.ServerInterface.
func (t *TaskHandler) PartiallyUpdateTask(ctx echo.Context, id int) error {
	panic("unimplemented")
}

// UpdateTask implements generated.ServerInterface.
func (t *TaskHandler) UpdateTask(ctx echo.Context, id int) error {
	panic("unimplemented")
}
