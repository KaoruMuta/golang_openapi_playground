package oapicodegen

import (
	"golang_openapi_playground/oapi_codegen/generated"
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
	var req generated.CreateTaskJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := t.r.CreateTask(req); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

// DeleteTask implements generated.ServerInterface.
func (t *TaskHandler) DeleteTask(ctx echo.Context, id int) error {
	if err := t.r.DeleteTask(id); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
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
func (t *TaskHandler) PartiallyUpdateTask(ctx echo.Context) error {
	var req generated.PartiallyUpdateTaskJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := t.r.PartiallyUpdateTask(req); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

// UpdateTask implements generated.ServerInterface.
func (t *TaskHandler) UpdateTask(ctx echo.Context) error {
	var req generated.UpdateTaskJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := t.r.UpdateTask(req); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
