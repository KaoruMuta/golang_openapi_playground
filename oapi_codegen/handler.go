package oapicodegen

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"golang_openapi_playground/model"
	"golang_openapi_playground/oapi_codegen/generated"
	"golang_openapi_playground/repository"
)

type TaskHandler struct {
	r repository.TaskRepository
}

func NewTaskHandler(r repository.TaskRepository) *TaskHandler {
	return &TaskHandler{r: r}
}

// CreateTask implements generated.ServerInterface.
func (t *TaskHandler) CreateTask(ctx echo.Context) error {
	var req generated.CreateTaskJSONRequestBody
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	task := model.Task{
		Title:     *req.Title,
		Completed: *req.Completed,
	}
	if err := t.r.CreateTask(task); err != nil {
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
	task := model.GetTaskFrom(req)
	if err := t.r.PartiallyUpdateTask(task); err != nil {
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
	task := model.GetTaskFrom(req)
	if err := t.r.UpdateTask(task); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
